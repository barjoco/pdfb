package main

import (
	"github.com/barjoco/pdfb"
)

func main() {
	pdf := pdfb.New()

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

	pdf.Page()

	pdf.SetFooter(
		"Arial",
		pdfb.TextAlign{Text: "Page {page} of {pages}", Align: "Centre"},
	)

	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Build PDF documents with ease.")

	pdf.Heading(1, "Heading 1")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(2, "Heading 2")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(3, "Heading 3")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(4, "Heading 4")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(5, "Heading 5")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Heading(6, "Heading 6")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud.")

	pdf.SaveAs("examples/hello/hello.pdf")
}
