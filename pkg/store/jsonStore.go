package store

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Andresch29/go-web/internal/domain"
)

type JsonStore struct {
	currentId int
	path string
}

func NewJsonStore(path string) (*JsonStore, error) {
	var products domain.Products
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &products)
	if err != nil {
		return nil, err
	}
	id := len(products)
	return &JsonStore{
		currentId: id,
		path: path,
	} , nil
}

func (j *JsonStore) Get() (domain.Products, error) {
	var products domain.Products
	content, err := ioutil.ReadFile(j.path)
	if err != nil {
		return products, err
	}

	err = json.Unmarshal(content, &products)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (j *JsonStore) Set(products domain.Products) error {
	data, err := json.MarshalIndent(products, "", "")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(j.path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (j *JsonStore) Create(product domain.Product) (domain.Product, error) {
	products, err := j.Get()
	if err != nil {
		return product, err
	}

	j.currentId++
	product.Id = j.currentId
	products = append(products, product)

	err = j.Set(products)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (j *JsonStore) GetAll() (domain.Products, error) {
	return j.Get()
}

func (j *JsonStore) GetById(id int) (*domain.Product, error) {
	products, err := j.Get()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		if p.Id == id {
			return &p, nil
		}
	}

	return nil, nil
}

func (j *JsonStore) GetByPrice(price float64) (domain.Products, error) {
	products, err := j.Get()
	if err != nil {
		return nil, err
	}

	result := domain.Products{}
	for _, p := range products {
		if p.Price > price {
			result = append(result, p)
		}
	}

	return result, nil
}

func (j *JsonStore) Update(product domain.Product) (*domain.Product, error) {
	var productDB *domain.Product
	products, err := j.Get()
	if err != nil {
		return productDB, err
	}

	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			productDB = &product

			err = j.Set(products)
			if err != nil {
				return productDB, err
			}

			return productDB, nil
		}
	}

	

	return productDB, nil
}

func (j *JsonStore) ExistsByCode(code string) (bool, error) {
	products, err := j.Get()
	if err != nil {
		return false, err
	}

	exists := false
	for _, p := range products {
		if p.CodeValue == code {
			exists = true
			return exists, nil
		}
	}

	return exists, nil
}

func (j *JsonStore) Delete(id int) (bool, error) {
	deleted := false
	products, err := j.Get()
	if err != nil {
		return deleted, err
	}

	var index int
	for i, p := range products {
		if p.Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return deleted, nil
	}
	products = append(products[:index], products[index+1:]...)

	err = j.Set(products)
	if err != nil {
		return deleted, err
	}
	return deleted, nil
}