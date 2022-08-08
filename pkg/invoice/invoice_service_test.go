package invoice

import (
	"testing"
)

type MockInvoiceRepository struct {
	getAllByFinancialYearCalls int
}

func (r *MockInvoiceRepository) GetByFinancialYear(y uint16) []Invoice {
	r.getAllByFinancialYearCalls++
	return []Invoice{
		{ID: 1, FinancialYear: y},
	}
}

func TestInvoiceService(t *testing.T) {
	r := MockInvoiceRepository{}
	s := NewInvoiceService(&r)

	t.Run("It should call repository when invokes GetAllByFinancialYear", func(t *testing.T) {
		s.GetAllByFinancialYear(2022)

		if r.getAllByFinancialYearCalls != 1 {
			t.Errorf("Number of calls to repository %v, expected %v", r.getAllByFinancialYearCalls, 1)
		}
	})

	t.Run("It should return one Invoice", func(t *testing.T) {
		invoices := s.GetAllByFinancialYear(2022)

		if len(invoices) != 1 {
			t.Errorf("Number of calls to repository %v, expected %v", r.getAllByFinancialYearCalls, 1)
		}
	})
}
