package invoice

type Service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return Service{
		repository: r,
	}
}

func (s Service) GetAllByFinancialYear(y uint16) []Invoice {
	return s.repository.GetByFinancialYear(y)
}
