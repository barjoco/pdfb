package pdfb

import (
	"encoding/base64"
	"strings"

	"github.com/barjoco/utils/log"
	"github.com/jung-kurt/gofpdf"
)

var err error
var pdf *gofpdf.Fpdf
var lineHeight float64
var fontSize float64

// Settings defines a list of document settings
type Settings struct {
	Orientation string
	Units       string
	PageSize    string
	Title       string
	Author      string
	Subject     string
	Keywords    []string
	Margin      float64
	LineHeight  float64
	FontSize    float64
}

// New initialises the PDF builder
func New(s Settings) {
	pdf = gofpdf.New(
		s.Orientation,
		s.Units,
		s.PageSize,
		"",
	)

	pdf.SetTitle(s.Title, true)
	pdf.SetAuthor(s.Author, true)
	pdf.SetSubject(s.Subject, true)
	pdf.SetKeywords(strings.Join(s.Keywords, " "), true)
	pdf.SetCreator("gofpdf", true)
	pdf.SetCellMargin(2)
	pdf.SetMargins(s.Margin-1, s.Margin, s.Margin-1)
	pdf.SetFontSize(s.FontSize)
	pdf.SetCreator("github.com/barjoco/pdfb", true)
	pdf.AliasNbPages("")
	PageBreak()

	lineHeight = s.LineHeight
	fontSize = s.FontSize

	// Import default fonts
	for _, f := range fontsBase64 {
		fontDecoded, err := base64.StdEncoding.DecodeString(f[1])
		log.ReportFatal(err)
		pdf.AddUTF8FontFromBytes(f[0], "", fontDecoded)
	}

	pdf.SetFont("Inter-Regular", "", 0)
}

// PageBreak is used to insert a new page
func PageBreak() {
	pdf.AddPage()
}

// SetLineHeight is used to set the line height
func SetLineHeight(lh float64) {
	if lh < 1 {
		log.ErrorFatal("Line height must be greater than 0")
	}
	lineHeight = lh
}

// GetLineHeight is used to get the line height
func GetLineHeight() float64 {
	return lineHeight
}

// Write is used to write text to the page
func Write(text string) {
	pdf.Write(lineHeight, text)
}

// SaveAs ...
func SaveAs(filePath string) {
	err := pdf.OutputFileAndClose(filePath)
	log.ReportFatal(err)
	log.Info("PDF saved to %s.", filePath)
}
