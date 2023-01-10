package product

import "github.com/Andresch29/go-web/internal/domain"

type Service interface {
	Create(product *domain.Product) (domain.Product, error)
	GetAll() domain.Products
	GetById(id int) (domain.Product, bool)
	GetByPrice(price float64) domain.Products
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(product *domain.Product) (domain.Product, error) {
	return s.repository.Create(product)
}

func (s *service) GetAll() domain.Products {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (domain.Product, bool) {
	return s.repository.GetById(id)
}

func (s *service) GetByPrice(price float64) domain.Products {
	return s.repository.GetByPrice(price)
}



