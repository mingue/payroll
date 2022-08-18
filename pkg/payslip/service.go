package payslip

import (
	"time"

	"github.com/mingue/payroll/pkg/invoice"
)

type Service struct {
	repository                 Repository
	invoiceRepository             invoice.Repository
	netSalaryCalculatorFactory NetSalaryCalculatorFactory
}

func NewService(
	r Repository,
	is invoice.Repository,
	f NetSalaryCalculatorFactory,
) Service {
	return Service{
		repository: r,
		invoiceRepository: is,
		netSalaryCalculatorFactory: f,
	}
}

func (s Service) CreatePayslip(paymentDate time.Time, countryIsoCode string) Payslip {
	var invoices = s.invoiceRepository.GetByFinancialYearAndCountryIsoCode(paymentDate.Year(), countryIsoCode)
	var existingPayslips = s.repository.GetByFinancialYearAndCountryIsoCode(paymentDate.Year(), countryIsoCode)

	var netCalculator = s.netSalaryCalculatorFactory.GetByDateAndCountryCode(paymentDate, countryIsoCode)
	payslip := netCalculator.Calculate(invoices, existingPayslips, paymentDate)

	s.repository.Save(&payslip)
	return payslip
}
