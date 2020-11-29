package pdfb

var stdFonts = []string{"courier", "helvetica", "arial", "times", "symbol", "zapfdingbats"}

// makes a font styleStr from the stored font style information
func (p *Pdfb) makeFontStyleStr() (styleStr string) {
	switch {
	case p.font.Bold:
		styleStr += "b"
	case p.font.Italic:
		styleStr += "i"
	case p.font.Underline:
		styleStr += "u"
	case p.font.Strikethrough:
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
	p.SetFontSize(font.Size)

	p.font = font

	p.pdf.SetFont(font.Family, p.makeFontStyleStr(), font.Size)
}

// SetFontSize is used to set the font size
func (p *Pdfb) SetFontSize(fontSize float64) {
	if fontSize == -1 {
		fontSize = 12
	}

	fontSizeIncrease := fontSize / p.font.Size

	// scale lineHeight with increase/decrease of fontSize
	p.lineHeight *= fontSizeIncrease

	p.font.Size = fontSize
	p.pdf.SetFontSize(fontSize)
}

// Bold is used to write in bold
func (p *Pdfb) Bold(str string) {
	p.font.Bold = true
	p.SetFont(p.font)

	p.Write(str)

	p.font.Bold = false
	p.SetFont(p.font)
}

// BoldLn is used to write in bold, then drop a line
func (p *Pdfb) BoldLn(str string) {
	p.Bold(str)
	p.Ln(1)
}
