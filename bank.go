package generator

import "github.com/jung-kurt/gofpdf"

type Bank struct {
	PayBy         string // Tax in percent ex 17
	BankName      string // Tax in amount ex 123.40
	Address       string
	AccountType   string
	IFSC          string
	SWIFT         string
	AccountNumber string
	pdf           *gofpdf.Fpdf
}

func (b *Bank) appendBankTODoc(x float64, y float64, fill bool, pdf *gofpdf.Fpdf) {
	b.pdf = pdf
	b.pdf.Ln(20)
	b.pdf.SetFont("Helvetica", "B", 10)
	b.text(40, 0, "Payment Details")
	b.pdf.Ln(5)
	b.pdf.SetFont("Helvetica", "", 10)
	headers := []string{
		"Pay By", "Bank Name", "Address", "Account Type (current/savings)",
		"A/c No.", "IFSC", "SWIFT (international)",
	}

	b.pdf.SetDrawColor(64, 64, 64)
	for i := 0; i < len(headers); i++ {
		b.pdf.SetFont("Helvetica", "B", 10)
		b.textFormat(60, 5, headers[i], "1", 0, "R", true, 0, "")
		b.pdf.SetFont("Helvetica", "", 10)
		switch i {
		case 0:
			b.textFormat(100, 5, b.PayBy, "1", 0, "L", false, 0, "")
			break
		case 1:
			b.textFormat(100, 5, b.BankName, "1", 0, "L", false, 0, "")
			break
		case 2:
			b.textFormat(100, 5, b.Address, "1", 0, "L", false, 0, "")
			break
		case 3:
			b.textFormat(100, 5, b.AccountType, "1", 0, "L", false, 0, "")
			break
		case 4:
			b.textFormat(100, 5, b.AccountNumber, "1", 0, "L", false, 0, "")
			break
		case 5:
			b.textFormat(100, 5, b.IFSC, "1", 0, "L", false, 0, "")
			break
		case 6:
			b.textFormat(100, 5, b.SWIFT, "1", 0, "L", false, 0, "")
			break
		}
		b.pdf.Ln(5)
	}
}

func (b *Bank) text(x, y float64, txtStr string) {
	unicodeToPDF := b.pdf.UnicodeTranslatorFromDescriptor("")
	b.pdf.Cell(x, y, unicodeToPDF(txtStr))
}

func (b *Bank) textFormat(w, h float64, txtStr string, borderStr string, ln int,
	alignStr string, fill bool, link int, linkStr string) {
	unicodeToPDF := b.pdf.UnicodeTranslatorFromDescriptor("")
	b.pdf.CellFormat(w, h, unicodeToPDF(txtStr), borderStr, ln, alignStr, fill, link, linkStr)
}
