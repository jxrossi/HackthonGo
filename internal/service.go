package internal

type Service interface {
	Store(table string, data interface{}) error
	ExistsCustomers() bool
	ExistsProducts() bool
	ExistsInvoices() bool
	ExistsSales() bool
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) Store(table string, data interface{}) error {
	return s.repository.Store(table, data)
}

func (s *service) ExistsCustomers() bool {
	return s.repository.ExistsCustomers()
}

func (s *service) ExistsProducts() bool {
	return s.repository.ExistsProducts()
}

func (s *service) ExistsInvoices() bool {
	return s.repository.ExistsInvoices()
}

func (s *service) ExistsSales() bool {
	return s.repository.ExistsSales()
}


