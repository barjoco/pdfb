package pdfb

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/barjoco/utils/colour"
	"github.com/barjoco/utils/log"
	"github.com/jung-kurt/gofpdf"
)

var err error

// Font defines a font
type Font struct {
	Family        string
	Size          float64
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
}

// TextAlign is used to define text with alignment
type TextAlign struct {
	Text  string
	Align string
}

// Pdfb ...
type Pdfb struct {
	pdf          *gofpdf.Fpdf
	bgFunc       func()
	orientation  string
	units        string
	pageSize     string
	title        string
	author       string
	subject      string
	keywords     []string
	margin       float64
	font         Font
	lineHeight   float64
	background   string
	foreground   string
	headerHeight float64
	footerHeight float64
}

// New returns a PDF Builder
func New() *Pdfb {
	// PDF default options
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
		20.0,
		Font{Family: "Arial", Size: 12.0},
		6.0,
		"#ffffff",
		"#000000",
		0,
		0,
	}

	// pdf initialisation
	p.pdf.SetTitle(p.title, true)
	p.pdf.SetAuthor(p.author, true)
	p.pdf.SetSubject(p.subject, true)
	p.pdf.SetKeywords(strings.Join(p.keywords, ";"), true)
	p.pdf.SetCreator("github.com/barjoco/pdfb", true)
	p.pdf.SetCellMargin(2)
	p.pdf.SetMargins(p.margin-1, p.margin, p.margin-1) //subtract half of cellMargin
	p.pdf.SetFontSize(p.font.Size)
	p.pdf.AliasNbPages("")
	p.pdf.SetAutoPageBreak(true, p.margin)
	p.pdf.SetFont(p.font.Family, "", p.font.Size)

	// bgFunc gets called in headerFunc, used to set the background
	// colour of the document
	p.bgFunc = func() {
		w, h, _ := p.pdf.PageSize(p.pdf.PageNo())
		currentR, currentG, currentB := p.pdf.GetFillColor()
		p.Box(0, 0, w, h, p.background, true, false)
		p.pdf.SetFillColor(currentR, currentG, currentB)
	}

	// default header, does nothing except set the background colour
	p.pdf.SetHeaderFunc(func() {
		p.bgFunc()
	})

	p.SetForeground(p.foreground)
	p.Page()

	return p
}

// used to generate an align string
func (p *Pdfb) makeAlignStr(alignInput string) (alignStr string) {
	alignInput = strings.ToLower(alignInput)

	switch {
	case alignInput == "l" || alignInput == "left":
		alignStr = "L"
	case alignInput == "c" || alignInput == "centre":
		alignStr = "C"
	case alignInput == "r" || alignInput == "right":
		alignStr = "R"
	default:
		log.ErrorFatal("Invalid align input (%s)", alignInput)
	}

	return
}

// SetMargin is used to set the margin
func (p *Pdfb) SetMargin(margin float64) {
	p.margin = margin
	p.pdf.SetMargins(margin, margin, margin)
}

// Page is used to insert a new page
func (p *Pdfb) Page() {
	p.pdf.AddPage()
}

// SetForeground is used to set the text colour
func (p *Pdfb) SetForeground(hex string) {
	r, g, b := colour.HexToRGB(hex)
	p.pdf.SetTextColor(r, g, b)
}

// SetHeader is used to set the header
func (p *Pdfb) SetHeader(fontFamily string, content ...TextAlign) {
	pageWidth, _, _ := p.pdf.PageSize(p.pdf.PageNo())
	sectionWidth := (pageWidth - p.margin*2) / float64(len(content))
	p.headerHeight = p.margin * 1.5

	p.pdf.SetHeaderFunc(func() {
		// copy the current font
		currentFont := p.fontCopy(p.font)

		// used to draw the background colour
		p.bgFunc()

		// put the header at the top of the page
		p.SetY(0)

		// set font for header text
		p.SetFont(Font{
			Family: fontFamily,
			Size:   12,
		})

		// create cells for each section
		for _, c := range content {
			p.pdf.CellFormat(sectionWidth, p.headerHeight, c.Text, "", 0, "M"+p.makeAlignStr(c.Align), false, 0, "")
		}

		// set the font back to how it was
		p.SetFont(currentFont)

		// set cursor to the bottom of the header
		p.pdf.SetX(p.margin)
		p.pdf.SetY(p.headerHeight)
	})
}

// SetFooter is used to set the footer
// The page number and number of pages can be used in the footer
// using {pages} and {pages}.
//
// Eg. "Page {page} of {pages}"
func (p *Pdfb) SetFooter(fontFamily string, content ...TextAlign) {
	pageWidth, pageHeight, _ := p.pdf.PageSize(p.pdf.PageNo())
	sectionWidth := (pageWidth - p.margin*2) / float64(len(content))
	p.footerHeight = p.margin * 1.5

	p.pdf.SetFooterFunc(func() {
		// copy the current font
		currentFont := p.fontCopy(p.font)

		// set cursor to the position where the top of the footer starts drawing
		p.SetY(pageHeight - p.footerHeight)

		// set font for header text
		p.SetFont(Font{
			Family: fontFamily,
			Size:   12,
		})

		// create cells for each section
		for _, c := range content {
			c.Text = strings.ReplaceAll(c.Text, "{page}", strconv.Itoa(p.pdf.PageNo()))
			c.Text = strings.ReplaceAll(c.Text, "{pages}", "{nb}")
			p.pdf.CellFormat(sectionWidth, p.footerHeight, c.Text, "", 0, "M"+p.makeAlignStr(c.Align), false, 0, "")
		}

		// set the font back to how it was
		p.SetFont(currentFont)
	})

	// set the space from the bottom where the auto page break gets triggered
	p.pdf.SetAutoPageBreak(true, p.footerHeight)
}

// SetX is used to set the cursor's horizontal position
func (p *Pdfb) SetX(x float64) {
	p.pdf.SetX(x)
}

// SetY is used to set the cursor's vertical position
func (p *Pdfb) SetY(y float64) {
	p.pdf.SetY(y)
}

// Box is used to draw a box
func (p *Pdfb) Box(x, y, w, h float64, hex string, fill, border bool) {
	var styleStr string

	if fill {
		styleStr += "F"
	}
	if border {
		styleStr += "D"
	}

	r, g, b := colour.HexToRGB(hex)
	p.pdf.SetFillColor(r, g, b)
	p.pdf.Rect(x, y, w, h, styleStr)
}

// Circle is used to draw a circle
func (p *Pdfb) Circle(x, y, radius float64, hex string, fill, border bool) {
	var styleStr string

	if fill {
		styleStr += "F"
	}
	if border {
		styleStr += "D"
	}

	r, g, b := colour.HexToRGB(hex)
	p.pdf.SetFillColor(r, g, b)
	p.pdf.Circle(x, y, radius, styleStr)
}

// SetLine is used to set the line colour and weight
func (p *Pdfb) SetLine(hex string, weight float64) {
	r, g, b := colour.HexToRGB(hex)
	p.pdf.SetDrawColor(r, g, b)
	p.pdf.SetLineWidth(weight)
}

// Ln is used to insert a new line (or multiple)
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

// Heading is used to write headings of various levels
func (p *Pdfb) Heading(level int, str string) {
	// copy current font
	currentFont := p.fontCopy(p.font)
	currentLH := p.lineHeight

	// set font and write content
	p.SetFont(Font{
		Family: p.font.Family,
		Bold:   true,
		Size:   12 * (1 + float64(7-level)*0.15), // 12 being the default font size
	})

	// check that the heading height + height of 1 line of regular text
	// can fit before the end of the page(-margin) or the footer if present
	// // get page height
	_, pageHeight, _ := p.pdf.PageSize(p.pdf.PageNo())
	// // choose whether bottomSpace should be end of page - margin
	// // or the height of the footer
	var bottomSpace float64
	if p.footerHeight > 0 {
		bottomSpace = p.footerHeight
	} else {
		bottomSpace = p.margin
	}
	// // do the check and print line if necessary
	if (p.pdf.GetY() + p.lineHeight + currentLH) > (pageHeight - bottomSpace) {
		p.Ln(1)
	}

	// write header
	p.WriteLn(str)

	// set font back to how it was
	p.SetFont(currentFont)
}

// Width gets the page width
func (p *Pdfb) Width() float64 {
	w, _ := p.pdf.GetPageSize()
	return w
}

// Height gets the page height
func (p *Pdfb) Height() float64 {
	_, h := p.pdf.GetPageSize()
	return h
}
