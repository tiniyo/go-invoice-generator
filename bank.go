package generator

import "github.com/jung-kurt/gofpdf"

type Bank struct {
	PayBy       string // Tax in percent ex 17
	BankName    string // Tax in amount ex 123.40
	Address     string
	AccountType string
	IFSC        string
	IBAN        string
	SWIFT       string
	pdf         *gofpdf.Fpdf
}

func (b *Bank) appendBankTODoc(x float64, y float64, fill bool, pdf *gofpdf.Fpdf) {
	pdf.SetXY(x, y)
	b.pdf = pdf
	pdf.Ln(20)
	pdf.SetFont("Helvetica", "B", 10)
	pdf.Cell(40, 0, "Payment Details")
	pdf.SetFont("Helvetica", "", 10)
	b.pdf.Ln(5)
}
