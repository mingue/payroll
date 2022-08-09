package invoice

import (
	"testing"
	i "github.com/mingue/payroll/pkg/invoice"
)

type MockRepository struct {
	getByFinancialYearCalls int
}

func (r *MockRepository) GetByFinancialYear(y uint16) []i.Invoice {
	r.getByFinancialYearCalls++
	return []i.Invoice{
		{ID: 1, FinancialYear: y},
	}
}

func TestService(t *testing.T) {
	r := MockRepository{}
	s := i.NewService(&r)

	t.Run("It should call repository when invokes GetByFinancialYear", func(t *testing.T) {
		s.GetAllByFinancialYear(2022)

		if r.getByFinancialYearCalls != 1 {
			t.Errorf("Number of calls to repository %v, expected %v", r.getByFinancialYearCalls, 1)
		}
	})

	t.Run("It should return one Invoice", func(t *testing.T) {
		invoices := s.GetAllByFinancialYear(2022)

		if len(invoices) != 1 {
			t.Errorf("Number of calls to repository %v, expected %v", r.getByFinancialYearCalls, 1)
		}
	})
}