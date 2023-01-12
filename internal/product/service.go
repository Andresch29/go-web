package product

import "github.com/Andresch29/go-web/internal/domain"

type Service interface {
	Create(product domain.Product) (domain.Product, error)
	GetAll() (domain.Products, error)
	GetById(id int) (*domain.Product, error)
	GetByPrice(price float64) (domain.Products, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(id int) (bool, error)
	ExistsByCode(code string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(product domain.Product) (domain.Product, error) {
	return s.repository.Create(product)
}

func (s *service) GetAll() (domain.Products, error) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (*domain.Product, error) {
	return s.repository.GetById(id)
}

func (s *service) GetByPrice(price float64) (domain.Products, error) {
	return s.repository.GetByPrice(price)
}

func (s *service) Update(product domain.Product) (*domain.Product, error) {
	return s.repository.Update(product)
}

func (s *service) Delete(id int) (bool, error) {
	return s.repository.Delete(id)
}

func (s *service) ExistsByCode(code string) (bool, error) {
	return s.repository.ExistsByCode(code)
}