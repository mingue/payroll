package repositories

import (
	i "github.com/mingue/payroll/pkg/invoice"
	"github.com/samber/lo"
)

var _ i.Repository = &InvoiceRepository{}

type InvoiceRepository struct {
	invoices []i.Invoice
}

func (r *InvoiceRepository) GetByFinancialYearAndCountryIsoCode(y int, countryIsoCode string) []i.Invoice {
	invoices := lo.Filter(r.invoices, func(v i.Invoice, _ int) bool {
		return v.FinancialYear == y && v.CountryIsoCode == countryIsoCode
	})

	return invoices
}
