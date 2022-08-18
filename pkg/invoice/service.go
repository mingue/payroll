package invoice

type Service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return Service{
		repository: r,
	}
}

func (s Service) GetByFinancialYearAndCountryIsoCode(y int, countryIsoCode string) []Invoice {
	return s.repository.GetByFinancialYearAndCountryIsoCode(y, countryIsoCode)
}
