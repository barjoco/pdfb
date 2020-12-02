package main

import (
	"github.com/barjoco/pdfb"
)

func main() {
	pdf := pdfb.New()

	pdf.Page()

	pdf.Circle(pdf.Width(), pdf.Height(), 150, "#fff5f5", true, false)
	pdf.Box(0, 0, pdf.Width(), 6, "#f00", true, false)

	pdf.SetY(80)

	pdf.SetFontSize(15)
	pdf.SetForeground("#f00")
	pdf.BoldLn("Demonstration of building PDFs using Pdfb")

	pdf.SetFontSize(40)
	pdf.SetForeground("#000")
	pdf.BoldLn("Here is an example")
	pdf.SetFontSize(-1)

	pdf.Ln(1)
	pdf.Box(pdf.GetX(), pdf.GetY(), 60, 6, "#f00", true, false)

	pdf.Ln(8)

	pdf.SetFontSize(-1)
	pdf.WriteLn("By John Smith")

	pdf.SetHeader(
		"Arial",
		pdfb.TextAlign{Text: "Left text", Align: "Left"},
		pdfb.TextAlign{Text: "Centre text", Align: "c"},
		pdfb.TextAlign{Text: "Right text", Align: "right"},
	)

	pdf.SetFooter(
		"Arial",
		pdfb.TextAlign{Text: "Page {page} of {pages}", Align: "Centre"},
	)

	pdf.ToC(1)

	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Build PDF documents with ease.")

	pdf.Heading(1, "Heading 1")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(2, "Heading 2")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(3, "Slightly longer Heading 3")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(4, "Heading 4")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(5, "Heading 5")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(6, "Heading 6")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	pdf.Heading(1, "Example of lists")

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
	pdf.Heading(1, "Images")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	pdf.Ln(1)
	pdf.Image("fish.png", "c", pdf.GetX(), pdf.GetY(), 0, 30)

	pdf.Ln(1)
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	pdf.SaveAs("examples/hello/hello.pdf")
}
