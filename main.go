package main

import (
	"fmt"
	"io/ioutil"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)

/* country specific sms and cost */
/* country specific voice calls and cost*/

func main() {
	doc, err := generator.New(generator.Invoice, &generator.Options{
		CurrencySymbol:    "INR ",
		CurrencyPrecision: 2,
		CurrencyDecimal:   ".",
		CurrencyThousand:  ",",
		TextTotalTax:      "GST",
	})

	doc.SetHeader(&generator.HeaderFooter{
		Text: "Tiniyo Inc.",
	})

	doc.SetFooter(&generator.HeaderFooter{
		Text:       "<center>Tiniyo Inc. Build the future of communications</center>",
		Pagination: true,
	})

	doc.SetRef("testref")

	doc.SetVersion("1.0.0")

	doc.SetDescription("SMS Voice billing")

	logoBytes, _ := ioutil.ReadFile("./logo.png")

	doc.SetCompany(&generator.Contact{
		Name: "Tiniyo Inc.",
		Logo: &logoBytes,
		Address: &generator.Address{
			Address:    "Founders Cube",
			Address2:   "Mahadevpura",
			PostalCode: "560048",
			City:       "Bangalore",
			Country:    "India",
		},
	})

	doc.SetCustomer(&generator.Contact{
		Name: "CricHeroes Pvt. Ltd",
		Address: &generator.Address{
			Address:    "Digicorp House, Ambawadi",
			PostalCode: "380015",
			City:       "Ahmedabad",
			Country:    "India",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     "[IND] Delivered SMS",
		UnitCost: "0.11",
		Quantity: "135000",
		Tax: &generator.Tax{
			Percent: "18",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     "[IND] Sent SMS",
		UnitCost: "0.11",
		Quantity: "55000",
		Tax: &generator.Tax{
			Percent: "18",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     "[IND] Failed SMS",
		UnitCost: "0.00",
		Quantity: "5000",
		Tax: &generator.Tax{
			Percent: "18",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     "Voice Calls",
		UnitCost: "0.00",
		Quantity: "0",
		Tax: &generator.Tax{
			Percent: "18",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     "OTP Verifications",
		UnitCost: "0.00",
		Quantity: "0",
		Tax: &generator.Tax{
			Percent: "18",
		},
	})

	pdf, err := doc.Build()

	if err != nil {
		fmt.Println(err.Error())
	}

	err = pdf.OutputFileAndClose("./out.pdf")

	if err != nil {
		fmt.Println(err.Error())
	}
}
