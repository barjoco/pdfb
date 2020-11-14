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

// Font ...
type Font struct {
	identifier     string
	fontDir        string
	regular        string
	bold           string
	italic         string
	otherVariables []FontVariable
}

// FontVariable ...
type FontVariable struct {
	name string
	path string
}

// DefineFont ...
// func (p *Pdfb) DefineFont(fonts []Font) {
// 	for font := range fonts {

// 	}
// }

// ImportFonts is used to import fonts (or a single font) for use inside the PDF builder.
// Unix-style file globbing can be used, e.g. "~/.fonts/Roboto_*.ttf"
func (p *Pdfb) ImportFonts(fontPaths ...string) {
	// store invalid fonts (wrong type, duplicate)
	var fontErrors [][]string

	color.Blue("Imported fonts:")

	// loop over all provided font paths
	for _, f := range fontPaths {
		f = strings.TrimSpace(f)
		f = expandPath(f)

		// get all files matching the glob pattern
		fontFiles, err := filepath.Glob(f)
		log.ReportFatal(err)

		// loop over files matching the glob
		for _, fontFile := range fontFiles {
			fileName := filepath.Base(fontFile)
			dirName := filepath.Dir(fontFile)
			id := strings.Split(fileName, ".ttf")[0]

			// check if a font with this identifier is imported already
			if _, exists := fonts[id]; !exists {
				// check the file is of the correct type
				buf, err := ioutil.ReadFile(fontFile)
				log.ReportFatal(err)
				kind, err := filetype.Match(buf)
				if kind.Extension == "ttf" {
					// add to fonts map
					fonts[id] = fontFile
					// add font into p.pdf
					p.pdf.SetFontLocation(dirName)
					p.pdf.AddUTF8Font(id, "", fileName)
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

// SetFont is used to set the current font for writing in the PDF document
//
// The fontIdentifier is the name of the font file without its extension.
// For example, the identifier for Roboto_Regular.ttf is Roboto_Regular
func (p *Pdfb) SetFont(fontIdentifier string, options ...string) {
	if len(options) > 4 {
		log.ErrorFatal("Too many options given to SetFont. Must have between 0 and 4.")
	}

	if strings.ToLower(fontIdentifier) == "default" {
		fontIdentifier = p.opts.FontFamily
	}

	var styleStr string

	for _, opt := range options {
		if opt == "bold" || "opt" == "b" {
			styleStr += "b"
			if !array.Contains(stdFonts, strings.ToLower(fontIdentifier)) {
				log.ErrorFatal("Can't use bold style with non-standard font")
			}
		} else if opt == "italic" || opt == "i" {
			styleStr += "i"
			if !array.Contains(stdFonts, strings.ToLower(fontIdentifier)) {
				log.ErrorFatal("Can't use italic style with non-standard font")
			}
		} else if opt == "underline" || opt == "u" {
			styleStr += "u"
		} else if opt == "strikethrough" || opt == "s" {
			styleStr += "s"
		} else {
			log.ErrorFatal("Invalid option (%s) given to SetFont.", opt)
		}
	}

	fontSize, _ := p.pdf.GetFontSize()
	p.pdf.SetFont(fontIdentifier, styleStr, fontSize)
}

// SetFontSize is used to set the font size
func (p *Pdfb) SetFontSize(fontSize float64) {
	p.pdf.SetFontSize(fontSize)
}

// ResetFont ...
func (p *Pdfb) ResetFont() {
	fontSize, _ := p.pdf.GetFontSize()
	p.pdf.SetFont(p.opts.FontFamily, "", fontSize)
}
