package product

import (
	"github.com/Andresch29/go-web/internal/domain"
)

type Repository interface {
	Create(product domain.Product) (domain.Product, error)
	GetAll() (domain.Products, error)
	GetById(id int) (*domain.Product, error)
	GetByPrice(price float64) (domain.Products, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(id int) (bool, error)
	ExistsByCode(code string) (bool, error)
}

type Memory struct {
	currentID int
	Products domain.Products
}

func NewMemory() *Memory {
	var products = domain.Products{}
	
	return &Memory{
		currentID: 0,
		Products: products,
	}
}

func (m *Memory) Create(product domain.Product) (domain.Product, error) {
	m.currentID++
	product.Id = m.currentID
	m.Products = append(m.Products, product)

	return product, nil
}

func (m *Memory) GetAll() (domain.Products, error) {
	return m.Products, nil
}

func (m *Memory) GetById(id int) (*domain.Product, error) {

	for _, p := range m.Products {
		if p.Id == id {
			return &p, nil
		}
	}

	return nil, nil
}

func (m *Memory) GetByPrice(price float64) (domain.Products, error) {
	products := domain.Products{}

	for _, p := range m.Products {
		if p.Price > price {
			products = append(products, p)
		}
	}

	return products, nil
}

func (m *Memory) Update(product domain.Product) (*domain.Product, error) {
	var productNew *domain.Product
	for i, p := range m.Products {
		if p.Id == product.Id {
			m.Products[i] = product
			productNew = &product
			return productNew, nil
		}
	}

	return productNew, nil
}

func (m *Memory) Delete(id int) (bool, error) {
	deleted := false
	var index int
	for i, p := range m.Products {
		if p.Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return deleted, nil
	}
	m.Products = append(m.Products[:index], m.Products[index+1:]...)
	return deleted, nil
}

func (m *Memory) ExistsByCode(code string) (bool, error) {
	exists := false
	for _, p := range m.Products {
		if p.CodeValue == code {
			exists = true
			return exists, nil
		}
	}

	return exists, nil
}