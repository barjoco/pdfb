package pdfb

import (
	"fmt"
	"strings"

	"github.com/barjoco/utils/log"
	"github.com/jung-kurt/gofpdf"
)

var err error

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
	p.pdf.SetKeywords(strings.Join(p.opts.Keywords, " "), true)
	p.pdf.SetCreator("github.com/barjoco/pdfb", true)
	p.pdf.SetCellMargin(2)
	p.pdf.SetMargins(p.opts.Margin-1, p.opts.Margin, p.opts.Margin-1) //subtract half of cellMargin
	p.pdf.SetFontSize(p.opts.FontSize)
	p.pdf.AliasNbPages("")

	p.pdf.SetFont(p.opts.FontFamily, "", p.opts.FontSize)

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
func (p *Pdfb) SetHeader() {
	pageWidth, _, _ := p.pdf.PageSize(p.pdf.PageNo())
	leftMargin, topMargin, rightMargin, _ := p.pdf.GetMargins()
	sectionWidth := (pageWidth - leftMargin - rightMargin) / 3
	headerHeight := topMargin + p.opts.LineHeight

	p.pdf.SetHeaderFunc(func() {
		// put the header at the top of the page
		p.pdf.SetY(0)

		// format header
		p.pdf.CellFormat(sectionWidth, headerHeight, "Left", "1", 0, "L", false, 0, "")
		p.pdf.CellFormat(sectionWidth, headerHeight, "Centre", "1", 0, "C", false, 0, "")
		p.pdf.CellFormat(sectionWidth, headerHeight, "Right", "1", 0, "R", false, 0, "")

		// set cursor to the height of the header + 1 new line
		p.pdf.SetX(leftMargin)
		p.pdf.SetY(headerHeight + p.opts.LineHeight)
	})
}

// Table is used to draw a table
func (p *Pdfb) Table() {
	text1 := "Officia ex veniam et cillum Lorem velit. Excepteur velit est dolore"
	text2 := " irure amet sit mollit labore. Officia enim sit proident aute veniam laboris "
	text3 := "id quis sit cupidatat dolore."
	p.Write(text1)
	p.SetFont("default", "bold")
	p.Write(text2)
	p.ResetFont()
	p.Write(text3)

	p.Ln(4)

	pageWidth, _, _ := p.pdf.PageSize(p.pdf.PageNo())
	leftMargin, _, rightMargin, _ := p.pdf.GetMargins()
	colWidth := (pageWidth - leftMargin - rightMargin) / 2

	var cursor float64

	text1split := p.pdf.SplitText(text1, colWidth)

	for i, line := range text1split {
		p.Write(line)
		cursor = p.pdf.GetX()
		if i < len(text1split)-1 {
			p.Ln(1)
		}
	}

	p.SetFont("default", "bold")

	text2split := p.pdf.SplitText(text2, colWidth-cursor)
	p.Write(text2split[0])
	p.Ln(1)

	text2 = strings.Join(text2split[1:], " ")

	for _, line := range p.pdf.SplitText(text2, colWidth) {
		p.WriteLn(line)
	}

	p.Ln(2)
	p.pdf.CellFormat(pageWidth-p.opts.Margin*2, p.opts.LineHeight, "|", "1", 0, "C", false, 0, "")

	p.ResetFont()

	for _, line := range p.pdf.SplitText(text2, colWidth) {
		p.WriteLn(line)
	}

	// use these:
	// p.pdf.SplitText
	// p.pdf.GetStringWidth
}
