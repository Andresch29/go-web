package db

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Andresch29/go-web/models"
)

func ReadFile(path string) (*models.Products, error) {
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