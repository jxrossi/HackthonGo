package internal

import "HacktonGo/internal/models"

type Service interface {
	Store(table string, data interface{}) error
	ExistsCustomers() bool
	ExistsProducts() bool
	ExistsInvoices() bool
	ExistsSales() bool
	ExistsTotalsInvoices() bool
	StoreTotalInvoices(id int) error
	Enunciado1() ([]models.Enunciado1, error)
	Enunciado2() ([]models.Enunciado2, error)
	Enunciado3() ([]models.Enunciado3, error)
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

func (s *service) ExistsTotalsInvoices() bool {
	return s.repository.ExistsTotalsInvoices()
}

func (s *service) StoreTotalInvoices(id int) error {
	return s.repository.StoreTotalInvoices(id)
}

func (s *service) Enunciado1() ([]models.Enunciado1, error) {
	return s.repository.Enunciado1()
}

func (s *service) Enunciado2() ([]models.Enunciado2, error) {
	return s.repository.Enunciado2()
}

func (s *service) Enunciado3() ([]models.Enunciado3, error) {
	return s.repository.Enunciado3()
}
