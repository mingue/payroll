package payslip

import (
	"time"

	"github.com/mingue/payroll/pkg/invoice"
)

// TODO: Make it an interface
type NetSalaryCalculatorFactory struct {
}

func (f NetSalaryCalculatorFactory) GetByDateAndCountryCode(date time.Time, countryCode string) NetSalaryCalculator{
	return projectedFixedRateCalculator{}
}

type NetSalaryCalculator interface {
	Calculate(invoices []invoice.Invoice, payslips []Payslip, paymentDate time.Time) Payslip
}

var _ NetSalaryCalculator = projectedFixedRateCalculator{}

type projectedFixedRateCalculator struct {
}

// Calculate implements NetSalaryCalculator
func (projectedFixedRateCalculator) Calculate(invoices []invoice.Invoice, payslips []Payslip, paymentDate time.Time) Payslip {
	// TODO: Implement calculation and add tests
	return Payslip{Id: 1}
}
