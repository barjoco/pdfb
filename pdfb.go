package pdfb

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/barjoco/utils/colour"
	"github.com/barjoco/utils/inter"
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

// heading is used to define a heading
type heading struct {
	text  string
	level int
	page  int
	link  int
}

// Pdfb is the main Pdfb struct
type Pdfb struct {
	pdf             *gofpdf.Fpdf
	bgFunc          func()
	orientation     string
	units           string
	pageSize        string
	title           string
	author          string
	subject         string
	keywords        []string
	margin          float64
	font            Font
	lineHeight      float64
	background      string
	foreground      string
	headerHeight    float64
	footerHeight    float64
	headings        []heading
	tocPage         int
	writingContents bool
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
		Font{Family: "Inter", Size: 12.0},
		6.0,
		"#ffffff",
		"#000000",
		0,
		0,
		[]heading{},
		-1,
		false,
	}

	// import inter to be used as the default font
	p.pdf.AddUTF8FontFromBytes("Inter", "", decode(inter.InterRegular))
	p.pdf.AddUTF8FontFromBytes("Inter", "b", decode(inter.InterBold))
	p.pdf.AddUTF8FontFromBytes("Inter", "i", decode(inter.InterItalic))
	p.pdf.AddUTF8FontFromBytes("Inter", "bi", decode(inter.InterBoldItalic))

	// pdf initialisation
	p.pdf.SetTitle(p.title, true)
	p.pdf.SetAuthor(p.author, true)
	p.pdf.SetSubject(p.subject, true)
	p.pdf.SetKeywords(strings.Join(p.keywords, ";"), true)
	p.pdf.SetCreator("github.com/barjoco/pdfb", true)
	p.pdf.SetMargins(p.margin, p.margin, p.margin)
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

	// set foreground
	p.SetForeground(p.foreground)

	return p
}

func decode(b64Str string) (b []byte) {
	b, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		log.ErrorFatal("%s", err)
	}
	return
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

// GetMargin is used to set the margin
func (p *Pdfb) GetMargin() float64 {
	margin, _, _, _ := p.pdf.GetMargins()
	return margin
}

// Page is used to insert a new page
func (p *Pdfb) Page() {
	p.pdf.AddPage()
}

// SetHeader is used to set the header
func (p *Pdfb) SetHeader(fontFamily string, content ...TextAlign) {
	pageWidth, _, _ := p.pdf.PageSize(p.pdf.PageNo())
	sectionWidth := (pageWidth - p.margin*2) / float64(len(content))
	p.headerHeight = 25.0

	p.pdf.SetHeaderFunc(func() {
		// copy the current font
		currentFont := p.fontCopy(p.font)

		// get current foreground
		currentFG := p.foreground

		// used to draw the background colour
		p.bgFunc()

		// put the header at the top of the page
		p.SetY(0)

		// set font for header text
		p.SetFont(Font{
			Family: fontFamily,
			Size:   12,
		})

		// set foreground
		p.SetForeground("#000")

		// create cells for each section
		for _, c := range content {
			p.pdf.CellFormat(sectionWidth, p.headerHeight, c.Text, "", 0, "M"+p.makeAlignStr(c.Align), false, 0, "")
		}

		// set the font back to how it was
		p.SetFont(currentFont)

		// set foreground back to how it was
		p.SetForeground(currentFG)

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
	p.footerHeight = 25.0

	triggeredPage := p.pdf.PageNo()

	p.pdf.SetFooterFunc(func() {
		// don't run on the page that SetFooter was called on
		if p.pdf.PageNo() == triggeredPage {
			return
		}

		// copy the current font
		currentFont := p.fontCopy(p.font)

		// get current foreground
		currentFG := p.foreground

		// set cursor to the position where the top of the footer starts drawing
		p.SetY(pageHeight - p.footerHeight)

		// set font for header text
		p.SetFont(Font{
			Family: fontFamily,
			Size:   12,
		})

		// set foreground
		p.SetForeground("#000")

		// create cells for each section
		for _, c := range content {
			// deal with offset caused by width of text being calculated
			// with the {page} and {pages} aliases (resulting text is shorter)
			var offset float64
			if strings.Contains(c.Text, "{page}") {
				c.Text = strings.ReplaceAll(c.Text, "{page}", strconv.Itoa(p.pdf.PageNo()))
				offset += p.pdf.GetStringWidth("{page}")
			}
			if strings.Contains(c.Text, "{pages}") {
				offset += p.pdf.GetStringWidth("{pages}")
			}
			offset /= 2
			offset -= p.pdf.GetStringWidth("00") / 2
			p.SetX(p.GetX() + offset)
			// print
			p.pdf.CellFormat(sectionWidth-offset, p.footerHeight, c.Text, "", 0, "M"+p.makeAlignStr(c.Align), false, 0, "")
		}

		// set the font back to how it was
		p.SetFont(currentFont)

		// set foreground back to how it was
		p.SetForeground(currentFG)
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

// GetX is used to set the cursor's horizontal position
func (p *Pdfb) GetX() float64 {
	return p.pdf.GetX()
}

// GetY is used to set the cursor's vertical position
func (p *Pdfb) GetY() float64 {
	return p.pdf.GetY()
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

// Line is used to draw lines from one point to another
func (p *Pdfb) Line(fromX, fromY, toX, toY float64, hex string, weight float64) {
	currentDrawR, currentDrawG, currentDrawB := p.pdf.GetDrawColor()
	currentWeight := p.pdf.GetLineWidth()

	p.pdf.SetDrawColor(colour.HexToRGB(hex))
	p.pdf.SetLineWidth(weight)
	p.pdf.Line(fromX, fromY, toX, toY)

	p.pdf.SetDrawColor(currentDrawR, currentDrawG, currentDrawB)
	p.pdf.SetLineWidth(currentWeight)
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
	text := fmt.Sprintf(format, a...)
	p.pdf.Write(p.lineHeight, text)
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

// SaveAs is used to save the PDF document to a file
func (p *Pdfb) SaveAs(filePath string) {
	p.pdf.RegisterAlias("{pages}", strconv.Itoa(p.pdf.PageCount()))

	// go back and write the ToC if necessary
	if p.tocPage > 0 {
		p.writingContents = true
		p.pdf.SetPage(p.tocPage)
		p.SetY(p.headerHeight)
		p.Heading(1, "Contents")
		p.lineHeight *= 1.5
		var headingsPerPage int
		var headingsPerPageSet bool
		// slice is capped at the end because 'Contents' itself
		// is a heading
		for i, heading := range p.headings[:len(p.headings)-1] {
			// handle overflows onto the next page
			if !headingsPerPageSet && (p.GetY()+p.lineHeight) > (p.Height()-p.footerHeight) {
				headingsPerPage = i
				headingsPerPageSet = true
			}
			if headingsPerPageSet && i%headingsPerPage == 0 {
				p.pdf.SetPage(p.pdf.PageNo() + 1)
				p.SetY(p.headerHeight)
			}

			// bold for headingText and headingPage
			p.font.Bold = true
			p.SetFont(p.font)

			// heading text
			headingTextWidth := p.pdf.GetStringWidth(heading.text)
			headingTextSpaces := strings.Repeat("    ", heading.level-1)
			headingTextSpacesWidth := p.pdf.GetStringWidth(headingTextSpaces)
			headingTextWithSpaces := headingTextSpaces + heading.text
			headingTextWithSpacesWidth := p.pdf.GetStringWidth(headingTextWithSpaces)

			// heading page
			headingPage := strconv.Itoa(heading.page)
			headingPageWidth := p.pdf.GetStringWidth(headingPage)

			// dots
			p.font.Bold = false
			p.SetFont(p.font)
			pageWidth := p.Width() - p.margin*2
			dotSpace := pageWidth - headingTextWithSpacesWidth - headingPageWidth
			var dots string
			for {
				if p.pdf.GetStringWidth(dots) >= dotSpace-p.pdf.GetStringWidth("...") {
					break
				}
				dots += "."
			}

			// print

			p.font.Bold = true
			p.SetFont(p.font)
			p.SetX(p.GetX() + headingTextSpacesWidth)
			p.pdf.CellFormat(headingTextWidth, p.lineHeight, heading.text, "", 0, "L", false, heading.link, "")

			p.font.Bold = false
			p.SetFont(p.font)
			p.Write(dots)

			p.font.Bold = true
			p.SetFont(p.font)
			p.pdf.WriteAligned(0, p.lineHeight, headingPage, "R")
			p.Ln(1)
		}
		p.pdf.SetPage(p.pdf.PageCount())
	}

	// output file
	fmt.Println("Saving PDF...")
	err := p.pdf.OutputFileAndClose(filePath)
	log.ReportFatal(err)
	log.Info("PDF saved to %s.", filePath)
}

// Heading is used to write headings of various levels
func (p *Pdfb) Heading(level int, str string) {
	// create heading link
	headingLink := p.pdf.AddLink()
	p.pdf.SetLink(headingLink, p.GetY(), p.pdf.PageNo())

	// add bookmark
	if !p.writingContents {
		p.pdf.Bookmark(str, level-1, -1)
	}

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
	// // choose whether bottomSpace should be end of page - margin
	// // or the height of the footer
	var bottomSpace float64
	if p.footerHeight > 0 {
		bottomSpace = p.footerHeight
	} else {
		bottomSpace = p.margin
	}
	// // do the check and print line if necessary
	// // description of the check:
	// // p.lineHeight = lineheight of the header
	// // currentLH = lineheight of the previous text (and future text)
	// // currentLH/4 = an extra quarter of a lineheight space for line
	if (p.pdf.GetY() + p.lineHeight + currentLH + currentLH/4) > (p.Height() - bottomSpace) {
		p.Ln(1)
	}

	// write header
	p.WriteLn(str)

	// set font back to how it was
	p.SetFont(currentFont)

	// quarter of a lineHeight of space below headings
	p.SetY(p.GetY() + p.lineHeight/4)

	// add heading to headings array
	p.headings = append(p.headings, heading{str, level, p.pdf.PageNo(), headingLink})
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

// ToC is used to generate table of contents from headings
func (p *Pdfb) ToC(numPages int) {
	// insert a new page, then go to the next page, leaving
	// a blank page for the ToC
	p.Page()
	p.tocPage = p.pdf.PageNo()
	for i := 0; i < numPages; i++ {
		p.Page()
	}
}

// ListItem defines an item to use in the List function
type ListItem struct {
	Level int
	Text  string
}

// List is used for writing lists
func (p *Pdfb) List(items []ListItem) {
	// copy current font
	currentFont := p.fontCopy(p.font)
	maxIndent := 10

	// loop through list items
	for _, item := range items {
		// indent in from margin (indents stop at level 8)
		if item.Level <= maxIndent {
			p.SetX(p.GetX() + float64(8*(item.Level)))
		} else {
			p.SetX(p.GetX() + float64(8*maxIndent))
		}

		// switch to symbol font
		p.font.Family = "zapfdingbats"
		p.SetFont(p.font)

		// switch case for bullet type
		if item.Level <= maxIndent {
			switch {
			case item.Level == 1 || item.Level%3 == 1:
				p.font.Size -= 5
				p.SetFont(p.font)
				p.Write("\x6c")
			case item.Level == 2 || item.Level%3 == 2:
				p.font.Size -= 6
				p.SetFont(p.font)
				p.Write("\x6d ")
			case item.Level == 3 || item.Level%3 == 0:
				p.font.Size -= 5
				p.SetFont(p.font)
				p.Write("\x6e")
			}
		} else {
			p.font.Size -= 5
			p.SetFont(p.font)
			p.Write("\x6c")
		}

		// small indent in from the bullet symbol
		p.SetX(p.GetX() + 5)

		// change back to current font
		p.font.Size = currentFont.Size
		p.SetFont(currentFont)

		// print
		p.pdf.MultiCell(0, p.lineHeight, item.Text, "", "", false)

		// move down a little bit after each list item
		p.SetY(p.GetY() + 1)
	}
}

// Image is used to insert an image
// Use 0 in place of w or h to keep the aspect ratio
func (p *Pdfb) Image(filename, align string, x, y, w, h float64) {
	// calc w and/or h values if 0 is given
	if w == 0 {
		info := p.pdf.RegisterImage("fish.png", "")
		w = h * info.Width() / info.Height()
	}
	if h == 0 {
		info := p.pdf.RegisterImage("fish.png", "")
		h = w * info.Height() / info.Width()
	}

	// align image for left, right, or centre
	align = strings.ToLower(align)
	switch {
	case align == "l" || align == "left" || align == "":
	case align == "c" || align == "centre":
		x = p.GetX() + (p.Width()-p.margin*2)/2 - w/2
	case align == "r" || align == "right":
		x = p.Width() - p.margin - w
	default:
		log.ErrorFatal("Invalid alignment supplied to Image (%s)", align)
	}

	// draw image
	p.pdf.Image(filename, x, y, w, h, true, "", 0, "")
}

// Debug is used for debugging purposes
func (p *Pdfb) Debug(str string) {
	fmt.Println(str)
}

// Hyperlink is used to print hyperlinks
func (p *Pdfb) Hyperlink(displayText, url string) {
	p.SetForeground("#00f")
	p.pdf.WriteLinkString(p.lineHeight, displayText, url)
	p.SetForeground("#000")
}
