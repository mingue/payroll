package invoice

import (
	"testing"
)

var _ Repository = &mockRepository{}

type mockRepository struct {
	getByFinancialYearCalls int
}

func (r *mockRepository) GetByFinancialYearAndCountryIsoCode(y int, countryIsoCode string) []Invoice {
	r.getByFinancialYearCalls++
	return []Invoice{
		{Id: 1, FinancialYear: y},
	}
}

func TestService(t *testing.T) {
	r := mockRepository{}
	s := NewService(&r)

	t.Run("It should call repository when invokes GetByFinancialYear", func(t *testing.T) {
		s.GetByFinancialYearAndCountryIsoCode(2022,"NZ")

		if r.getByFinancialYearCalls != 1 {
			t.Errorf("Number of calls to repository %v, expected %v", r.getByFinancialYearCalls, 1)
		}
	})

	t.Run("It should return one Invoice", func(t *testing.T) {
		invoices := s.GetByFinancialYearAndCountryIsoCode(2022, "NZ")

		if len(invoices) != 1 {
			t.Errorf("Number of invoices returned %v, expected %v", r.getByFinancialYearCalls, 1)
		}
	})
}
