package pdfb

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/barjoco/utils/colour"
	"github.com/barjoco/utils/inter"
	"github.com/barjoco/utils/log"
	"github.com/jung-kurt/gofpdf"
)

var err error

// Pdfb is the main Pdfb struct
type Pdfb struct {
	pdf *gofpdf.Fpdf

	bgFunc          func()
	footerHeight    float64
	headerHeight    float64
	headings        []heading
	tocPage         int
	writingContents bool

	// customisable
	accentColour     string
	author           string
	background       string
	creationDate     time.Time
	font             Font
	foreground       string
	keywords         []string
	lineHeight       float64
	margin           float64
	modificationDate time.Time
	orientation      string
	pageHeight       float64
	pageSize         string
	pageWidth        float64
	subject          string
	title            string
}

// New returns a PDF Builder
func New() *Pdfb {
	// PDF default options
	p := &Pdfb{
		pdf: gofpdf.New("P", "mm", "A4", ""),

		bgFunc:          func() {},
		footerHeight:    0,
		headerHeight:    0,
		headings:        []heading{},
		tocPage:         -1,
		writingContents: false,

		accentColour:     "#f00",
		author:           "Author",
		background:       "#ffffff",
		creationDate:     time.Now(),
		font:             Font{Family: "Inter", Size: 12.0},
		foreground:       "#000000",
		keywords:         []string{},
		lineHeight:       6.0,
		margin:           20.0,
		modificationDate: time.Now(),
		orientation:      "P",
		pageHeight:       297.0,
		pageSize:         "A4",
		pageWidth:        210.0,
		subject:          "Creating documents with Pdfb",
		title:            "Document",
	}

	// import inter to be used as the default font
	p.pdf.AddUTF8FontFromBytes("Inter", "", decode(inter.InterRegular))
	p.pdf.AddUTF8FontFromBytes("Inter", "b", decode(inter.InterBold))
	p.pdf.AddUTF8FontFromBytes("Inter", "i", decode(inter.InterItalic))
	p.pdf.AddUTF8FontFromBytes("Inter", "bi", decode(inter.InterBoldItalic))

	p.checkpoint("Inter font imported")

	// pdf initialisation
	p.pdf.SetProducer("GoFPDF 2.17.2", true)
	p.pdf.AliasNbPages("")
	p.pdf.SetAuthor(p.author, true)
	p.pdf.SetAutoPageBreak(true, p.margin)
	p.pdf.SetCreator("github.com/barjoco/pdfb", true)
	p.pdf.SetCreationDate(p.creationDate)
	p.pdf.SetFont(p.font.Family, "", p.font.Size)
	p.pdf.SetFontSize(p.font.Size)
	p.pdf.SetKeywords(strings.Join(p.keywords, ";"), true)
	p.pdf.SetMargins(p.margin, p.margin, p.margin)
	p.pdf.SetModificationDate(p.modificationDate)
	p.pdf.SetSubject(p.subject, true)
	p.pdf.SetTextColor(colour.HexToRGB(p.foreground))
	p.pdf.SetTitle(p.title, true)

	p.checkpoint("PDF initialised")

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

	return p
}

// SetAccentColour is used to set the accentColour
func (p *Pdfb) SetAccentColour(accentColour string) {
	p.accentColour = accentColour
}

// GetAccentColour is used to get the accentColour
func (p *Pdfb) GetAccentColour() string {
	return p.accentColour
}

// SetAuthor is used to set the author
func (p *Pdfb) SetAuthor(author string) {
	p.author = author
	p.pdf.SetAuthor(author, true)
}

// GetAuthor is used to get the author
func (p *Pdfb) GetAuthor() string {
	return p.author
}

// SetBackground is used to set the background
func (p *Pdfb) SetBackground(background string) {
	p.background = background
}

// GetBackground is used to get the background
func (p *Pdfb) GetBackground() string {
	return p.background
}

// SetCreationDate is used to set the creationDate
func (p *Pdfb) SetCreationDate(creationDate time.Time) {
	p.creationDate = creationDate
	p.pdf.SetCreationDate(creationDate)
}

// GetCreationDate is used to get the creationDate
func (p *Pdfb) GetCreationDate() time.Time {
	return p.creationDate
}

// SetKeywords is used to set the keywords
func (p *Pdfb) SetKeywords(keywords []string) {
	p.keywords = keywords
	p.pdf.SetKeywords(strings.Join(keywords, ";"), true)
}

// GetKeywords is used to get the keywords
func (p *Pdfb) GetKeywords() []string {
	return p.keywords
}

// SetLineHeight is used to set the lineHeight
func (p *Pdfb) SetLineHeight(lineHeight float64) {
	p.lineHeight = lineHeight
}

// GetLineHeight is used to get the lineHeight
func (p *Pdfb) GetLineHeight() float64 {
	return p.lineHeight
}

// SetMargin is used to set the margin
func (p *Pdfb) SetMargin(margin float64) {
	p.margin = margin
	p.pdf.SetMargins(margin, margin, margin)
}

// GetMargin is used to get the margin
func (p *Pdfb) GetMargin() float64 {
	return p.margin
}

// SetModificationDate is used to set the modificationDate
func (p *Pdfb) SetModificationDate(modificationDate time.Time) {
	p.modificationDate = modificationDate
	p.pdf.SetModificationDate(modificationDate)
}

// GetModificationDate is used to get the modificationDate
func (p *Pdfb) GetModificationDate() time.Time {
	return p.modificationDate
}

// SetOrientation is used to set the orientation
func (p *Pdfb) SetOrientation(orientation string) {
	p.orientation = orientation
}

// GetOrientation is used to get the orientation
func (p *Pdfb) GetOrientation() string {
	return p.orientation
}

// SetPageHeight is used to set the pageHeight
func (p *Pdfb) SetPageHeight(pageHeight float64) {
	p.pageHeight = pageHeight
}

// GetPageHeight is used to get the pageHeight
func (p *Pdfb) GetPageHeight() float64 {
	_, h := p.pdf.GetPageSize()
	return h
}

// SetPageSize is used to set the pageSize
func (p *Pdfb) SetPageSize(pageSize string) {
	p.pageSize = pageSize
	switch strings.ToLower(pageSize) {
	case "a1":
		p.SetPageHeight(841.0)
		p.SetPageWidth(594.0)
		p.pageSize = "A1"
	case "a2":
		p.SetPageHeight(594.0)
		p.SetPageWidth(420.0)
		p.pageSize = "A2"
	case "a3":
		p.SetPageHeight(420.0)
		p.SetPageWidth(297.0)
		p.pageSize = "A3"
	case "a4":
		p.SetPageHeight(297.0)
		p.SetPageWidth(210.0)
		p.pageSize = "A4"
	case "a5":
		p.SetPageHeight(210.0)
		p.SetPageWidth(148.0)
		p.pageSize = "A5"
	case "a6":
		p.SetPageHeight(148.0)
		p.SetPageWidth(105.0)
		p.pageSize = "A6"
	case "letter":
		p.SetPageHeight(279.4)
		p.SetPageWidth(215.9)
		p.pageSize = "Letter"
	case "legal":
		p.SetPageHeight(355.6)
		p.SetPageWidth(215.9)
		p.pageSize = "Legal"
	case "tabloid":
		p.SetPageHeight(431.8)
		p.SetPageWidth(279.4)
		p.pageSize = "Tabloid"
	default:
		log.ErrorFatal("%s is not a valid page size.", pageSize)
	}
}

// GetPageSize is used to get the pageSize
func (p *Pdfb) GetPageSize() string {
	return p.pageSize
}

// SetPageWidth is used to set the pageWidth
func (p *Pdfb) SetPageWidth(pageWidth float64) {
	p.pageWidth = pageWidth
}

// GetPageWidth is used to get the pageWidth
func (p *Pdfb) GetPageWidth() float64 {
	w, _ := p.pdf.GetPageSize()
	return w
}

// SetSubject is used to set the subject
func (p *Pdfb) SetSubject(subject string) {
	p.subject = subject
	p.pdf.SetSubject(subject, true)
}

// GetSubject is used to get the subject
func (p *Pdfb) GetSubject() string {
	return p.subject
}

// SetTitle is used to set the title
func (p *Pdfb) SetTitle(title string) {
	p.title = title
	p.pdf.SetTitle(title, true)
}

// GetTitle is used to get the title
func (p *Pdfb) GetTitle() string {
	return p.title
}

//
//	End of setters and getters
//

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

// Page is used to insert a new page
func (p *Pdfb) Page() {
	p.pdf.AddPageFormat(p.GetOrientation(), gofpdf.SizeType{
		Wd: p.GetPageWidth(),
		Ht: p.GetPageHeight(),
	})
	p.checkpoint("Page added")
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
	p.checkpoint("Y position set")
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

	p.pdf.SetFillColor(colour.HexToRGB(hex))
	p.pdf.Rect(x, y, w, h, styleStr)

	p.checkpoint("Box created")
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

	p.pdf.SetFillColor(colour.HexToRGB(hex))
	p.pdf.Circle(x, y, radius, styleStr)

	p.checkpoint("Circle created")
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
	p.pdf.SetDrawColor(colour.HexToRGB(hex))
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
	p.finalFunc()

	// output file
	fmt.Println("Saving PDF...")
	err := p.pdf.OutputFileAndClose(filePath)
	log.ReportFatal(err)
	log.Info("PDF saved to %s.", filePath)
}

// Heading is used to write headings of various levels
func (p *Pdfb) Heading(level int, str string) {
	// level must be 1-6
	if level < 1 || level > 6 {
		log.ErrorFatal("Invalid level supplied to Heading (%d)", level)
	}

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
	currentForeground := p.foreground

	// pick font size
	var fontSize float64
	switch level {
	case 6:
		fontSize = 12
	case 5:
		fontSize = 12.5
	case 4:
		fontSize = 13.5
	case 3:
		fontSize = 15
	case 2:
		fontSize = 17
	case 1:
		fontSize = 19.5
	}

	// set font and write content
	p.SetFont(Font{
		Family: p.font.Family,
		Bold:   true,
		Size:   fontSize,
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
	if (p.pdf.GetY() + p.lineHeight + currentLH + currentLH/4) > (p.GetPageHeight() - bottomSpace) {
		p.Ln(1)
	}

	// set foreground for header level 1
	if level == 1 {
		p.SetForeground(p.accentColour)
	}

	// write header
	p.WriteLn(str)

	// draw line under for header level 1
	if level == 1 {
		p.Line(p.margin, p.GetY(), p.GetPageWidth()-p.margin, p.GetY(), p.accentColour, 0.5)
		p.SetY(p.GetY() + p.lineHeight*0.25) // larger gap below heading due to line
	} else {
		p.SetY(p.GetY() + p.lineHeight*0.1) // gap below heading
	}

	// set font back to how it was
	p.SetFont(currentFont)
	p.SetLineHeight(currentLH)
	p.SetForeground(currentForeground)

	// add heading to headings array
	p.headings = append(p.headings, heading{str, level, p.pdf.PageNo(), headingLink})

	p.checkpoint("Heading created")
}

// ToC is used to generate table of contents from headings
func (p *Pdfb) ToC(numPages int) {
	// insert number of pages for toc
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
			p.SetX(p.GetX() + float64(6*(item.Level)))
		} else {
			p.SetX(p.GetX() + float64(6*maxIndent))
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
		p.SetX(p.GetX() + 4)

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
		x = p.GetX() + (p.GetPageWidth()-p.margin*2)/2 - w/2
	case align == "r" || align == "right":
		x = p.GetPageWidth() - p.margin - w
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
	currentFG := p.GetForeground()

	p.SetForeground("#00f")
	p.pdf.WriteLinkString(p.lineHeight, displayText, url)

	p.SetForeground(currentFG)
}

// This is the function that gets called before any "outputting" methods
// such as SaveAs or ExportAs
func (p *Pdfb) finalFunc() {
	p.pdf.RegisterAlias("{pages}", strconv.Itoa(p.pdf.PageCount()))

	// go back and write the ToC if necessary
	if p.tocPage > 0 {
		// needed to prevent writing a bookmark for 'contents' heading
		p.writingContents = true
		// go to toc page
		p.pdf.SetPage(p.tocPage)
		// go to top of page, under heading or margin
		if p.headerHeight > 0 {
			p.SetY(p.headerHeight)
		} else {
			p.SetY(p.margin)
		}
		p.Heading(1, "Contents")
		p.lineHeight *= 1.5
		var headingsPerPage int
		var headingsPerPageSet bool
		// slice is capped at the end because 'Contents' itself
		// is a heading
		for i, heading := range p.headings[:len(p.headings)-1] {
			// handle overflows onto the next page
			if !headingsPerPageSet && (p.GetY()+p.lineHeight) > (p.GetPageHeight()-p.footerHeight) {
				headingsPerPage = i
				headingsPerPageSet = true
			}
			if headingsPerPageSet && i%headingsPerPage == 0 {
				p.pdf.SetPage(p.pdf.PageNo() + 1)
				p.SetY(p.headerHeight)
			}

			// set font to bold for level 1 headings
			if heading.level == 1 {
				p.font.Bold = true
				p.SetFont(p.font)
			} else {
				p.font.Bold = false
				p.SetFont(p.font)
			}

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
			pageWidth := p.GetPageWidth() - p.margin*2
			dotSpace := pageWidth - headingTextWithSpacesWidth - headingPageWidth
			var dots string
			for {
				if p.pdf.GetStringWidth(dots) >= dotSpace-p.pdf.GetStringWidth("...") {
					break
				}
				dots += "."
			}

			//
			//	print
			//

			if heading.level == 1 {
				p.font.Bold = true
				p.SetFont(p.font)
			}

			// heading text
			p.SetX(p.GetX() + headingTextSpacesWidth)
			p.pdf.CellFormat(headingTextWidth, p.lineHeight, heading.text, "", 0, "L", false, heading.link, "")

			// dots
			p.SetX(p.GetX() + 0.25) // move to the right a little bit
			p.Write(dots)

			// heading page number
			p.pdf.WriteAligned(0, p.lineHeight, headingPage, "R")
			p.Ln(1)
		}

		// go back to the end of the document before output
		p.pdf.SetPage(p.pdf.PageCount())
	}
}

// ExportAsBase64 is used to return a base64 encoding of the PDF
func (p *Pdfb) ExportAsBase64() string {
	p.finalFunc()
	buf := new(bytes.Buffer)
	err := p.pdf.Output(buf)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

// Error is used to print the PDF error
func (p *Pdfb) Error() {
	log.Info("PDF Error: %s", p.pdf.Error())
}
