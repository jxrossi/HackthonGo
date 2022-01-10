package internal

type Service interface {
	StoreCustomers(table string, data interface{}) error
	ExistsCustomers() bool
	ExistsProducts() bool
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) StoreCustomers(table string, data interface{}) error {
	err := s.repository.StoreCustomers(table, data); if err != nil {
		return err
	}

	return nil
}

func (s *service) ExistsCustomers() bool {
	return s.repository.ExistsCustomers()
}

func (s *service) ExistsProducts() bool {
	return s.repository.ExistsProducts()
}


