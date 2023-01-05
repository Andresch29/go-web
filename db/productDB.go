package db

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Andresch29/go-web/models"
)

func NewDB(path string) (*models.Products, error) {
	var products = &models.Products{}
	content, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

	err = json.Unmarshal(content, products)
	if err != nil {
        return products, err
    }

	return products, nil
}

func GetAllProducts() (*models.Products, error) {
	productsDB, err := NewDB("/Users/andhenao/Documents/bootcamp/goweb/products.json")
	if err != nil {
		return productsDB, err
	}
	return productsDB, nil
}

func GetProductById(id int) (models.Product, bool, error) {
	var product = models.Product{}
	productsDB, err := NewDB("/Users/andhenao/Documents/bootcamp/goweb/products.json")
	if err != nil {
		return product, false, err
	}

	ok := false
	for _, p := range *productsDB {
		if p.Id == id {
			product = p
			ok = true
			break
		}
	}

	return product, ok, nil
}

func GetProductByPrice(price float64) (models.Products, error) {
	products := models.Products{}
	productsDB, err := NewDB("/Users/andhenao/Documents/bootcamp/goweb/products.json")
	if err != nil {
		return products, err
	}

	for _, p := range *productsDB {
		if p.Price > price {
			products = append(products, p)
		}
	}

	return products, nil
}
