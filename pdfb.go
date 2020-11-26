package pdfb

import (
	"fmt"
	"strings"

	"github.com/barjoco/utils/colour"
	"github.com/barjoco/utils/log"
	"github.com/jung-kurt/gofpdf"
)

var err error

// Pdfb ...
type Pdfb struct {
	pdf         *gofpdf.Fpdf
	bgFunc      func()
	orientation string
	units       string
	pageSize    string
	title       string
	author      string
	subject     string
	keywords    []string
	margin      float64
	fontFamily  string
	fontSize    float64
	lineHeight  float64
	background  string
	foreground  string
}

// New returns a PDF Builder
func New() *Pdfb {
	p := &Pdfb{
		gofpdf.New("P", "mm", "A4", ""),
		func() {},
		"P",
		"mm",
		"A4",
		"Document",
		"Author",
		"",
		[]string{},
		10.0,
		"Arial",
		12.0,
		5.0,
		"#ffffff",
		"#000000",
	}

	// pdf initialisation
	p.pdf.SetTitle(p.title, true)
	p.pdf.SetAuthor(p.author, true)
	p.pdf.SetSubject(p.subject, true)
	p.pdf.SetKeywords(strings.Join(p.keywords, ";"), true)
	p.pdf.SetCreator("github.com/barjoco/pdfb", true)
	p.pdf.SetCellMargin(2)
	p.pdf.SetMargins(p.margin-1, p.margin, p.margin-1) //subtract half of cellMargin
	p.pdf.SetFontSize(p.fontSize)
	p.pdf.AliasNbPages("")
	p.pdf.SetAutoPageBreak(true, p.margin)
	p.pdf.SetFont(p.fontFamily, "", p.fontSize)

	p.bgFunc = func() {
		w, h, _ := p.pdf.PageSize(p.pdf.PageNo())
		currentR, currentG, currentB := p.pdf.GetFillColor()
		p.Box(0, 0, w, h, p.background)
		p.pdf.SetFillColor(currentR, currentG, currentB)
	}

	p.pdf.SetHeaderFunc(func() {
		p.bgFunc()
	})

	p.SetForeground(p.foreground)

	return p
}

// SetForeground ...
func (p *Pdfb) SetForeground(hex string) {
	fgRGB, err := colour.HexToRGB(hex)
	log.ReportFatal(err)
	p.pdf.SetTextColor(int(fgRGB.R), int(fgRGB.G), int(fgRGB.B))
}

// Box ...
func (p *Pdfb) Box(x, y, w, h float64, hex string) {
	RGB, err := colour.HexToRGB(hex)
	log.ReportFatal(err)
	p.pdf.SetFillColor(int(RGB.R), int(RGB.G), int(RGB.B))
	p.pdf.Rect(x, y, w, h, "F")
}

// Page is used to insert a new page
func (p *Pdfb) Page() {
	p.pdf.AddPage()
}

// Ln is used to insert a new line
func (p *Pdfb) Ln(lines int) {
	for i := 0; i < lines; i++ {
		p.pdf.Ln(p.lineHeight)
	}
}

// Write is used to write text to the page
func (p *Pdfb) Write(format string, a ...interface{}) {
	p.pdf.Write(p.lineHeight, fmt.Sprintf(format, a...))
}

// WriteLn is used to write text to the page (drop to next line after text)
func (p *Pdfb) WriteLn(format string, a ...interface{}) {
	p.Write(format, a...)
	p.Ln(1)
}

// Paragraph is used to write a paragraph (blank line after text)
func (p *Pdfb) Paragraph(format string, a ...interface{}) {
	p.Write(format, a...)
	p.Ln(2)
}

// SaveAs ...
func (p *Pdfb) SaveAs(filePath string) {
	err := p.pdf.OutputFileAndClose(filePath)
	log.ReportFatal(err)
	log.Info("PDF saved to %s.", filePath)
}

// SetHeader ...
// The height of the header is the height of the top margin + the line height
func (p *Pdfb) SetHeader(leftText, centreText, rightText string) {
	pageWidth, _, _ := p.pdf.PageSize(p.pdf.PageNo())
	margin, _, _, _ := p.pdf.GetMargins()
	sectionWidth := (pageWidth - margin*2) / 3
	headerHeight := margin + p.lineHeight

	p.pdf.SetHeaderFunc(func() {
		// bgfunc is necessary, used for drawing the background colour rectangle
		p.bgFunc()

		// put the header at the top of the page
		p.pdf.SetY(0)

		// format header
		p.pdf.CellFormat(sectionWidth, headerHeight, leftText, "", 0, "L", false, 0, "")
		p.pdf.CellFormat(sectionWidth, headerHeight, centreText, "", 0, "C", false, 0, "")
		p.pdf.CellFormat(sectionWidth, headerHeight, rightText, "", 0, "R", false, 0, "")

		// set cursor to the height of the header + 1 new line
		p.pdf.SetX(margin)
		p.pdf.SetY(headerHeight + p.lineHeight)
	})
}

// SetFooter ...
func (p *Pdfb) SetFooter() {
	pageWidth, pageHeight, _ := p.pdf.PageSize(p.pdf.PageNo())
	margin, _, _, _ := p.pdf.GetMargins()
	sectionWidth := (pageWidth - margin*2)
	footerHeight := margin + p.lineHeight

	p.pdf.SetFooterFunc(func() {
		// set cursor to the position where the top of the footer starts drawing
		p.pdf.SetY(pageHeight - footerHeight)

		// format footer
		str := fmt.Sprintf("Page %d of {nb}", p.pdf.PageNo())
		p.pdf.CellFormat(sectionWidth, footerHeight, str, "", 0, "C", false, 0, "")
	})

	// set the space from the bottom where the auto page break gets triggered
	p.pdf.SetAutoPageBreak(true, footerHeight+p.lineHeight*0.75)
}
