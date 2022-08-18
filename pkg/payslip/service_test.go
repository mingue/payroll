package payslip

import (
	"testing"
	"time"

	"github.com/mingue/payroll/pkg/invoice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {
	r := new(MockRepository)
	ir := new(InvoiceMockRepository)
	f := NetSalaryCalculatorFactory{}
	s := NewService(r, ir, f)

	t.Run("CreatePayslip should save a Payslip", func(t *testing.T) {
		now := time.Now()
		countryCode := "NZ"

		ir.On("GetByFinancialYearAndCountryIsoCode", now.Year(), countryCode).
			Return([]invoice.Invoice{
				{Id: 1},
			})

		r.On("GetByFinancialYearAndCountryIsoCode", now.Year(), countryCode).
			Return(Payslip{Id: 1})

		r.On("Save", &Payslip{Id: 1})

		payslip := s.CreatePayslip(now, countryCode)

		assert.NotNil(t, payslip)
		assert.Equal(t, 1, payslip.Id)
		ir.AssertExpectations(t)
		r.AssertExpectations(t)
	})
}

var _ Repository = &MockRepository{}

type MockRepository struct {
	mock.Mock
}

// GetByFinancialYearAndCountryIsoCode implements Repository
func (m *MockRepository) GetByFinancialYearAndCountryIsoCode(year int, countryIsoCode string) []Payslip {
	m.Called(year, countryIsoCode)

	return []Payslip{
		{CountryIsoCode: countryIsoCode, PaymentDate: time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
}

// Save implements Repository
func (m *MockRepository) Save(p *Payslip) {
	m.Called(p)
}

var _ invoice.Repository = &InvoiceMockRepository{}

type InvoiceMockRepository struct {
	mock.Mock
}

// GetByFinancialYearAndCountryIsoCode implements invoice.Repository
func (m *InvoiceMockRepository) GetByFinancialYearAndCountryIsoCode(y int, countryIsoCode string) []invoice.Invoice {
	m.Called(y, countryIsoCode)

	return []invoice.Invoice{
		{},
	}
}
