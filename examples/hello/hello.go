package main

import (
	"github.com/barjoco/pdfb"
)

func main() {
	pdfb.New(pdfb.Settings{
		Orientation: "P",
		Units:       "mm",
		PageSize:    "A4",
		Title:       "PDF Doc",
		Author:      "John Smith",
		Subject:     "Creating PDFs with pdfb",
		Keywords:    []string{"pdf", "builder"},
		Margin:      10.0,
		LineHeight:  8.0,
		FontSize:    14.0,
	})

	pdfb.ImportFonts("~/.fonts/Roboto_Mono/*")

	pdfb.Write("Hello ")
	pdfb.SetFont("RobotoMono-Medium")
	pdfb.Write("and welcome to ")
	pdfb.StandardFont("times", "bus")
	pdfb.Write("PDF Builder.")
	pdfb.SaveAs("examples/hello/hello.pdf")
}
