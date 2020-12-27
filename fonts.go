package pdfb

import (
	"path"
	"strings"

	"github.com/barjoco/utils/colour"
	"github.com/barjoco/utils/log"
)

var stdFonts = []string{"courier", "helvetica", "arial", "times", "symbol", "zapfdingbats"}

// Font defines a font
type Font struct {
	Family        string
	Size          float64
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
}

// makes a font styleStr from the stored font style information
func (p *Pdfb) makeFontStyleStr() (styleStr string) {
	if p.font.Bold {
		styleStr += "b"
	}
	if p.font.Italic {
		styleStr += "i"
	}
	if p.font.Underline {
		styleStr += "u"
	}
	if p.font.Strikethrough {
		styleStr += "s"
	}

	return
}

// creates a copy of a font with the same font properties
func (p *Pdfb) fontCopy(f Font) Font {
	return Font{
		Family:        f.Family,
		Size:          f.Size,
		Bold:          f.Bold,
		Italic:        f.Italic,
		Underline:     f.Underline,
		Strikethrough: f.Strikethrough,
	}
}

// SetFont is used to set the font
func (p *Pdfb) SetFont(font Font) {
	if font.Size == 0 {
		font.Size = p.font.Size
	}
	if font.Family == "" {
		font.Family = p.font.Family
	}
	if strings.ToLower(font.Family) == "default" {
		font.Family = "Inter"
	}

	// call this before settings the p.font, since SetFontSize uses a comparison
	// of the old p.lineHeight, so p.font need not be overwritten before setting the new fontSize
	p.SetFontSize(font.Size)

	// call this after SetFontSize to set the new p.font
	p.font = font

	// set font within pdf
	p.pdf.SetFont(font.Family, p.makeFontStyleStr(), font.Size)

	// check for errors re: fonts
	if p.pdf.Err() {
		log.ErrorFatal("%s", p.pdf.Error())
	}
}

// GetFont is used to get the font
func (p *Pdfb) GetFont(font Font) Font {
	return p.font
}

// SetFontSize is used to set the font size
func (p *Pdfb) SetFontSize(fontSize float64) {
	// scale lineHeight with increase/decrease of fontSize
	p.lineHeight *= fontSize / p.font.Size

	// set new fontSize
	p.font.Size = fontSize
	p.pdf.SetFontSize(fontSize)

	p.checkpoint("Font sized changed")
}

// FontStyle defines a custom font style
// eg. FontStyle{"RobotoMono-Bold.ttf", "Bold"}
type FontStyle struct {
	File  string
	Style string
}

// ImportFont is used to import custom fonts
func (p *Pdfb) ImportFont(fontName, fontDir string, fontStyles []FontStyle) {
	for _, fontStyle := range fontStyles {
		style := strings.ToLower(fontStyle.Style)
		var styleStr string

		switch {
		case style == "" || style == "regular":
		case style == "b" || style == "bold":
			styleStr += "b"
		case style == "i" || style == "italic":
			styleStr += "i"
		case style == "bi" || style == "bolditalic":
			styleStr += "bi"
		default:
			log.ErrorFatal("Invalid font style supplied to ImportFont (%s)", fontStyle.Style)
		}

		p.pdf.AddUTF8Font(fontName, styleStr, path.Join(fontDir, fontStyle.File))
	}
}

// SetForeground is used to set the text colour
func (p *Pdfb) SetForeground(hex string) {
	p.foreground = hex
	p.pdf.SetTextColor(colour.HexToRGB(hex))

	p.checkpoint("Foreground set")
}

// GetForeground is used to get the foreground
func (p *Pdfb) GetForeground() string {
	return p.foreground
}

// Bold is used to print bold text
func (p *Pdfb) Bold(str string) {
	p.font.Bold = true
	p.SetFont(p.font)

	p.Write(str)

	p.font.Bold = false
	p.SetFont(p.font)

	p.checkpoint("Bold text written")
}

// BoldLn is used to print bold text, then print new line
func (p *Pdfb) BoldLn(str string) {
	p.Bold(str)
	p.Ln(1)
}

// Italic is used to print italic text
func (p *Pdfb) Italic(str string) {
	p.font.Italic = true
	p.SetFont(p.font)

	p.Write(str)

	p.font.Italic = false
	p.SetFont(p.font)
}

// ItalicLn is used to print italic text, then print new line
func (p *Pdfb) ItalicLn(str string) {
	p.Italic(str)
	p.Ln(1)
}

// BoldItalic is used to print bold italic text
func (p *Pdfb) BoldItalic(str string) {
	p.font.Bold = true
	p.font.Italic = true
	p.SetFont(p.font)

	p.Write(str)

	p.font.Bold = false
	p.font.Italic = false
	p.SetFont(p.font)
}

// BoldItalicLn is used to print bold italic text, then print new line
func (p *Pdfb) BoldItalicLn(str string) {
	p.BoldItalic(str)
	p.Ln(1)
}
