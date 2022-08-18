package payslip

type Repository interface {
	GetByFinancialYearAndCountryIsoCode(year int, countryIsoCode string) []Payslip
	Save(p *Payslip)
}