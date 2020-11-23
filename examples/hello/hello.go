package main

import (
	"github.com/barjoco/pdfb"
)

func main() {
	opts := &pdfb.Options{
		Orientation: "P",
		Units:       "mm",
		PageSize:    "A4",
		Title:       "PDF Doc",
		Author:      "John Smith",
		Subject:     "Creating PDFs with pdfb",
		Keywords:    []string{"pdf", "builder"},
		Margin:      10.0,
		FontFamily:  "Default",
		FontSize:    12.0,
		LineHeight:  5.0,
		Foreground:  "#BCBDD0",
		Background:  "#161923",
	}

	pdf := pdfb.New(opts)

	pdf.DefineFonts(
		pdfb.Font{
			Identifier: "RobotoMono",
			FontDir:    "~/.fonts/Roboto_Mono",
			Styles: []pdfb.FontStyle{
				{Name: "Regular", File: "RobotoMono-Regular"},
				{Name: "Bold", File: "RobotoMono-Bold"},
				{Name: "Italic", File: "RobotoMono-Italic"},
				{Name: "BoldItalic", File: "RobotoMono-BoldItalic"},
				{Name: "Thin", File: "RobotoMono-Thin"},
			},
		},
	)

	pdf.SetHeader("Left", "Centre", "Right")
	pdf.SetFooter()

	pdf.Page()
	pdf.SetFont("RobotoMono", "BoldItalic", "Strikethrough")
	pdf.Write("Hello %s ", opts.Author)
	pdf.SetFont("RobotoMono", "Thin", "underline")
	pdf.WriteLn("and")
	pdf.Write("welcome to ")
	pdf.SetFont("Times")
	pdf.Write("PDF Builder.")
	pdf.Ln(2)
	pdf.ResetFont()
	pdf.Write("Build PDF documents with ease.")
	pdf.Ln(2)
	pdf.Paragraph("Exercitation mollit veniam velit ex aliquip occaecat commodo Lorem fugiat. Occaecat voluptate Lorem sint consequat consequat incididunt consectetur elit aliqua id. Culpa dolor irure culpa sint cupidatat aliqua sint excepteur laborum. Aliqua ea cupidatat ut irure officia in proident incididunt exercitation anim amet. Ea deserunt ex Lorem consequat labore Lorem deserunt consequat ad aute cupidatat Lorem. Tempor voluptate quis consequat exercitation est ex qui dolore est consectetur est deserunt ut nostrud. Laborum consectetur exercitation nostrud exercitation anim culpa ullamco mollit aute nulla consectetur incididunt est tempor. Non enim non pariatur irure incididunt reprehenderit culpa in consectetur non cupidatat ea in nisi. Ea magna sunt minim non et qui sunt excepteur elit laborum id. Voluptate quis aliquip eu proident occaecat laborum nostrud nostrud sint ea eiusmod Lorem fugiat. Do culpa ut deserunt id aliquip sunt Lorem. Incididunt excepteur dolore est magna tempor enim quis cillum nisi quis laboris.")
	pdf.Paragraph("Officia ex veniam et cillum Lorem velit. Excepteur velit est dolore commodo irure amet sit mollit labore. Officia enim sit proident aute veniam laboris id id quis sit cupidatat dolore. Lorem dolore tempor est anim ea aliqua in aliquip sunt incididunt veniam pariatur enim. Qui sint excepteur quis Lorem cillum voluptate eu duis.")
	// pdf.Table()
	pdf.Page()
	pdf.Write("New page")
	pdf.SaveAs("examples/hello/hello.pdf")
}
