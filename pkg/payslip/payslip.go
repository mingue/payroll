package payslip

import "time"

type Payslip struct {
	Id             int
	PaymentDate    time.Time
	Amount         float32
	CountryIsoCode string
}
