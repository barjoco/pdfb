package main

import "github.com/barjoco/pdfb"

func main() {
	pdf := pdfb.New()

	pdf.Page()

	pdf.SaveAs("examples/tests/tests.pdf")
}
