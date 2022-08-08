package invoice

type InvoiceService struct {
	repository IInvoiceRepository
}

func NewInvoiceService(r IInvoiceRepository) InvoiceService {
	return InvoiceService{
		repository: r,
	}
}

func (s *InvoiceService) GetAllByFinancialYear(y uint16) []Invoice {
	return s.repository.GetByFinancialYear(y)
}
