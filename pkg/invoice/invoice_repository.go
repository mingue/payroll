package invoice

type IInvoiceRepository interface {
	GetByFinancialYear(y uint16) []Invoice
}
type InvoiceRepository struct {
}

func (r *InvoiceRepository) GetByFinancialYear(y uint16) []Invoice {
	return []Invoice{
		{ID: 1, FinancialYear: y},
		{ID: 1, FinancialYear: y},
	}
}
