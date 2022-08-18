package repositories

import (
	p "github.com/mingue/payroll/pkg/payslip"
	"github.com/samber/lo"
)

var _ p.Repository = &PayslipRepository{}

type PayslipRepository struct {
	payslips []p.Payslip
}

// Save implements payslip.Repository
func (r *PayslipRepository) Save(p *p.Payslip) {
	p.Id = len(r.payslips)
	r.payslips = append(r.payslips, *p)
}

func (r *PayslipRepository) GetByFinancialYearAndCountryIsoCode(year int, countryIsoCode string) []p.Payslip {
	payslips := lo.Filter(r.payslips, func(v p.Payslip, _ int) bool {
		return v.PaymentDate.Year() == year && v.CountryIsoCode == countryIsoCode
	})

	return payslips
}
