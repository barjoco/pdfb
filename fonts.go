package pdfb

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/barjoco/utils/array"

	"github.com/barjoco/utils/log"
	"github.com/fatih/color"
	"github.com/h2non/filetype"
)

var stdFonts = []string{"courier", "helvetica", "arial", "times", "symbol", "zapfdingbats"}
var fonts = make(map[string]string)

// ImportFonts is used to import fonts (or a single font) for use inside the PDF builder.
// Unix-style file globbing can be used, e.g. "~/.fonts/Roboto_*.ttf"
func ImportFonts(fontPaths ...string) {
	// store invalid fonts (wrong type, duplicate)
	var fontErrors [][]string

	color.Blue("Imported fonts:")

	for _, f := range fontPaths {
		f = strings.TrimSpace(f)
		f = expandPath(f)

		// get all files matching the glob pattern
		fontFiles, err := filepath.Glob(f)
		log.ReportFatal(err)

		for _, fontFile := range fontFiles {
			// get file name before the .ttf, to be used as the font identifier
			fileName := filepath.Base(fontFile)
			dirName := filepath.Dir(fontFile)
			id := strings.Split(fileName, ".ttf")[0]

			// check if a font with this identifier is imported already
			if _, exists := fonts[id]; !exists {
				// check the file is of the correct type
				buf, _ := ioutil.ReadFile(fontFile)
				kind, _ := filetype.Match(buf)
				if kind.Extension == "ttf" {
					// add to fonts map
					fonts[id] = fontFile
					pdf.SetFontLocation(dirName)
					pdf.AddUTF8Font(id, "", fileName)
					fmt.Println(color.CyanString(id+":"), color.GreenString("✔"), color.HiBlackString("("+fontFile+")"))
				} else {
					fontErrors = append(fontErrors, []string{id, fontFile, "Invalid type. Must be font type TTF"})
				}
			} else {
				fontErrors = append(fontErrors, []string{id, fontFile, "A font with this name has already been imported"})
			}
		}
	}

	// print out any errors that were caught
	for _, e := range fontErrors {
		fmt.Println(color.RedString(e[0]+":"), color.HiRedString(e[2]), color.HiBlackString("("+e[1]+")"))
	}
}

// GetFonts is used to see which fonts have been exported, and their identifiers
func GetFonts() map[string]string {
	return fonts
}

// SetFont is used to set the current font for writing in the PDF document
//
// The fontIdentifier is the name of the font file without its extension.
// For example, the identifier for Roboto_Regular.ttf is Roboto_Regular
func SetFont(fontIdentifier string) {
	pdf.SetFont(fontIdentifier, "", 0)
}

// SetFontSize is used to set the font size
func SetFontSize(s float64) {
	pdf.SetFontSize(s)
	fontSize = s
}

// StandardFont ...
func StandardFont(fontName, fontStyle string) {
	fontName = strings.ToLower(fontName)

	if array.Contains(stdFonts, fontName) {
		pdf.SetFont(fontName, fontStyle, fontSize)
	} else {
		log.ErrorFatal("%s is not a standard font.", fontName)
	}
}
