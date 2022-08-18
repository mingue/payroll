package invoice

type Repository interface {
	GetByFinancialYearAndCountryIsoCode(y int, countryIsoCode string) []Invoice
}
