package main

import (
	"github.com/barjoco/pdfb"
)

func main() {
	pdf := pdfb.New()

	// PDF Metadata

	pdf.SetAuthor("John Smith")
	pdf.SetTitle("Pdfb Document")
	pdf.SetSubject("Building documents with Pdfb")
	pdf.SetKeywords([]string{"Pdfb", "Document", "Example"})

	//
	//	Front page
	//

	pdf.Page()

	pdf.Circle(pdf.GetPageWidth(), pdf.GetPageHeight(), 150, "#fff5f5", true, false)
	pdf.Box(0, 0, pdf.GetPageWidth(), 6, pdf.GetAccentColour(), true, false)

	pdf.SetY(80)

	pdf.SetFontSize(15)
	pdf.SetForeground(pdf.GetAccentColour())
	pdf.BoldLn("Demonstration of building PDFs using Pdfb")
	pdf.SetY(pdf.GetY() + 4)

	pdf.SetFontSize(40)
	pdf.SetForeground("#000")
	pdf.BoldLn("Here is an example")
	pdf.SetY(pdf.GetY() + 6)

	pdf.BoxInline(60, 6, pdf.GetAccentColour(), true, false)

	pdf.SetFontSize(15)
	pdf.Ln(6)
	pdf.BoldLn("John Smith")

	pdf.SetFontSize(12)

	//
	//	Headers and footers
	//

	pdf.SetHeader(
		"Inter",
		pdfb.TextAlign{Text: "Left text", Align: "Left"},
		pdfb.TextAlign{Text: "Centre text", Align: "c"},
		pdfb.TextAlign{Text: "Right text", Align: "right"},
	)

	pdf.SetFooter(
		"Inter",
		pdfb.TextAlign{Text: "Page {page} of {pages}", Align: "Centre"},
	)

	//
	//	Table of Contents
	//

	pdf.Page()
	pdf.ToC(1)

	//
	//	Headings
	//

	pdf.Heading(1, "Heading 1")
	pdf.Heading(2, "Heading 2")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(3, "Heading 3")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(4, "Heading 4")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(5, "Heading 5")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(6, "Heading 6")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	//
	//	Lists
	//

	pdf.Heading(1, "Lists")

	pdf.List(
		[]pdfb.ListItem{
			{Level: 1, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 1, Text: "1 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 2, Text: "2 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 3, Text: "3 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 4, Text: "4 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 5, Text: "5 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 6, Text: "6 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 7, Text: "7 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 8, Text: "8 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 9, Text: "9 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 10, Text: "10 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 11, Text: "11 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 12, Text: "12 Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 2, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 1, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 2, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
		},
	)
	pdf.Ln(1)

	//
	//	Images
	//

	pdf.Heading(1, "Images")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	pdf.Image("./examples/hello/fish.png", "c", pdf.GetX(), pdf.GetY(), 0, 70)

	pdf.Ln(1)
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	//
	//	Custom fonts
	//

	pdf.Heading(1, "Custom fonts")

	pdf.ImportFont("RobotoMono", "./examples/hello/RobotoMono",
		[]pdfb.FontStyle{
			{File: "RobotoMono-Regular.ttf", Style: ""},
			{File: "RobotoMono-Bold.ttf", Style: "Bold"},
			{File: "RobotoMono-Italic.ttf", Style: "Italic"},
			{File: "RobotoMono-BoldItalic.ttf", Style: "BoldItalic"},
		},
	)

	pdf.SetFont(pdfb.Font{Family: "RobotoMono"})
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	pdf.SetFont(pdfb.Font{Family: "Inter"})

	pdf.Write("Here is some ")
	pdf.BoldLn("bold text.")

	pdf.Write("Here is some ")
	pdf.ItalicLn("italic text.")

	pdf.Write("Here is some ")
	pdf.BoldItalicLn("bold italic text.")

	pdf.Write("Here is some ")
	pdf.SetFont(pdfb.Font{Bold: true, Underline: true})
	pdf.WriteLn("bold underline text.")
	pdf.SetFont(pdfb.Font{})

	pdf.Write("Here is some ")
	pdf.SetFont(pdfb.Font{Italic: true, Strikethrough: true})
	pdf.WriteLn("italic strikethrough text.")
	pdf.SetFont(pdfb.Font{})

	pdf.Ln(1)

	pdf.SetFont(pdfb.Font{Family: "RobotoMono"})
	pdf.Heading(2, "Heading with custom font")
	pdf.WriteLn("List with custom font:")
	pdf.List(
		[]pdfb.ListItem{
			{Level: 1, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 2, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 3, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 2, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
			{Level: 1, Text: "Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat."},
		},
	)

	pdf.SetFont(pdfb.Font{Family: "Default"})
	pdf.Ln(1)

	pdf.Heading(1, "Hyperlinks")

	pdf.Write("Here is a ")
	pdf.Hyperlink("hyperlink", "https://github.com/barjoco/pdfb")
	pdf.WriteLn(" to the Pdfb repo.")

	pdf.SaveAs("examples/hello/hello.pdf")
}
