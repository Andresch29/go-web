package product

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/Andresch29/go-web/internal/domain"
)

type Repository interface {
	Create(product *domain.Product) (domain.Product, error)
	GetAll() domain.Products
	GetById(id int) (domain.Product, bool)
	GetByPrice(price float64) domain.Products
}

type Memory struct {
	currentID int
	Products domain.Products
}

func NewMemory() *Memory {
	var products = domain.Products{}
	content, _ := ioutil.ReadFile("/Users/andhenao/Documents/bootcamp/goweb/products.json")

	json.Unmarshal(content, &products)
	
	currentID := len(products)

	return &Memory{
		currentID: currentID,
		Products: products,
	}
}

func (m *Memory) Create(product *domain.Product) (domain.Product, error) {
	if product == nil {
		return *product, errors.New("Product can't be nil")
	}

	m.currentID++
	product.Id = m.currentID
	m.Products = append(m.Products, *product)

	return *product, nil
}

func (m *Memory) GetAll() domain.Products {
	return m.Products
}

func (m *Memory) GetById(id int) (domain.Product, bool) {
	var product domain.Product

	ok := false
	for _, p := range m.Products {
		if p.Id == id {
			product = p
			ok = true
			break
		}
	}

	return product, ok
}

func (m *Memory) GetByPrice(price float64) domain.Products {
	products := domain.Products{}

	for _, p := range m.Products {
		if p.Price > price {
			products = append(products, p)
		}
	}

	return products
}