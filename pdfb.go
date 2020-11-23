package pdfb

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/barjoco/utils/colour"
	"github.com/barjoco/utils/log"
	"github.com/jung-kurt/gofpdf"
)

var err error
var fgColour color.RGBA
var bgColour color.RGBA
var bgFunc func(*Pdfb)

// Options defines a list of document options
type Options struct {
	Orientation string
	Units       string
	PageSize    string
	Title       string
	Author      string
	Subject     string
	Keywords    []string
	Margin      float64
	FontFamily  string
	FontSize    float64
	LineHeight  float64
	Background  string
	Foreground  string
}

// Pdfb ...
type Pdfb struct {
	pdf  *gofpdf.Fpdf
	opts *Options
}

// New returns a PDF Builder
func New(opts *Options) *Pdfb {
	// Default font is Arial
	if strings.ToLower(opts.FontFamily) == "default" {
		opts.FontFamily = "arial"
	}

	// New pdfb instance
	p := &Pdfb{
		gofpdf.New(
			opts.Orientation,
			opts.Units,
			opts.PageSize,
			"",
		),
		opts,
	}

	// pdf initialisation
	p.pdf.SetTitle(p.opts.Title, true)
	p.pdf.SetAuthor(p.opts.Author, true)
	p.pdf.SetSubject(p.opts.Subject, true)
	p.pdf.SetKeywords(strings.Join(p.opts.Keywords, ";"), true)
	p.pdf.SetCreator("github.com/barjoco/pdfb", true)
	p.pdf.SetCellMargin(2)
	p.pdf.SetMargins(p.opts.Margin-1, p.opts.Margin, p.opts.Margin-1) //subtract half of cellMargin
	p.pdf.SetFontSize(p.opts.FontSize)
	p.pdf.AliasNbPages("")
	p.pdf.SetAutoPageBreak(true, p.opts.Margin)
	p.pdf.SetFont(p.opts.FontFamily, "", p.opts.FontSize)

	// set foreground colour
	fgColour, err = colour.HexToRGB(p.opts.Foreground)
	log.ReportFatal(err)
	p.pdf.SetTextColor(int(fgColour.R), int(fgColour.G), int(fgColour.B))

	// set background colour
	bgColour, err = colour.HexToRGB(p.opts.Background)
	log.ReportFatal(err)

	// set background colour
	bgFunc = func(p *Pdfb) {
		oldR, oldG, oldB := p.pdf.GetFillColor()
		w, h, _ := p.pdf.PageSize(p.pdf.PageNo())
		p.pdf.SetFillColor(int(bgColour.R), int(bgColour.G), int(bgColour.B))
		p.pdf.Rect(0, 0, w, h, "F")
		p.pdf.SetFillColor(oldR, oldG, oldB)
	}

	// default headerfunc draws a rectangle under each page
	// filled with opts.Background colour
	p.pdf.SetHeaderFunc(func() {
		bgFunc(p)
	})

	return p
}

// Page is used to insert a new page
func (p *Pdfb) Page() {
	p.pdf.AddPage()
}

// Ln is used to insert a new line
func (p *Pdfb) Ln(lines int) {
	for i := 0; i < lines; i++ {
		p.pdf.Ln(p.opts.LineHeight)
	}
}

// Write is used to write text to the page
func (p *Pdfb) Write(format string, a ...interface{}) {
	p.pdf.Write(p.opts.LineHeight, fmt.Sprintf(format, a...))
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
	headerHeight := margin + p.opts.LineHeight

	p.pdf.SetHeaderFunc(func() {
		// bgfunc used to draw rectangle for background colour
		bgFunc(p)

		// put the header at the top of the page
		p.pdf.SetY(0)

		// format header
		p.pdf.CellFormat(sectionWidth, headerHeight, leftText, "", 0, "L", false, 0, "")
		p.pdf.CellFormat(sectionWidth, headerHeight, centreText, "", 0, "C", false, 0, "")
		p.pdf.CellFormat(sectionWidth, headerHeight, rightText, "", 0, "R", false, 0, "")

		// set cursor to the height of the header + 1 new line
		p.pdf.SetX(margin)
		p.pdf.SetY(headerHeight + p.opts.LineHeight)
	})
}

// SetFooter ...
func (p *Pdfb) SetFooter() {
	pageWidth, pageHeight, _ := p.pdf.PageSize(p.pdf.PageNo())
	margin, _, _, _ := p.pdf.GetMargins()
	sectionWidth := (pageWidth - margin*2)
	footerHeight := margin + p.opts.LineHeight

	p.pdf.SetFooterFunc(func() {
		p.pdf.SetY(pageHeight - footerHeight)
		pageNo := p.pdf.PageNo()

		str := fmt.Sprintf("Page %d of {nb}", pageNo)
		p.pdf.CellFormat(sectionWidth, footerHeight, str, "", 0, "C", false, 0, "")
	})

	// set the space from the bottom where the auto page break gets triggered
	p.pdf.SetAutoPageBreak(true, footerHeight+p.opts.LineHeight*0.75)
}
