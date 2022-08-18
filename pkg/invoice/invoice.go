package invoice

import "time"

type Invoice struct {
	Id            int
	InvoiceDate   time.Time
	PaymentDate   time.Time
	Amount        float32
	FinancialYear int
	GST           float32
	CountryIsoCode    string
}
