package generator

import (
	"github.com/shopspring/decimal"
)

type Discount struct {
	Percent string // Tax in percent ex 17
	Amount  string // Tax in amount ex 123.40
}

func (d *Discount) getDiscount() (string, decimal.Decimal) {
	var discount string
	var discountType string = "percent"

	if len(d.Percent) > 0 {
		discount = d.Percent
	}

	if len(d.Amount) > 0 {
		discount = d.Amount
		discountType = "amount"
	}

	decVal, _ := decimal.NewFromString(discount)

	return discountType, decVal
}
