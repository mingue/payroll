package invoice

type Repository interface {
	GetByFinancialYear(y uint16) []Invoice
}
