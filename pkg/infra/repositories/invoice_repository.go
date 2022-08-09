package invoice

import (
	i "github.com/mingue/payroll/pkg/invoice"
)

type InvoiceRepository struct {}

func (r InvoiceRepository) GetByFinancialYear(y uint16) []i.Invoice {
	return []i.Invoice{
		{ID: 1, FinancialYear: y},
		{ID: 1, FinancialYear: y},
	}
}
