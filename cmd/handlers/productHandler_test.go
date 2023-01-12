package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Andresch29/go-web/internal/product"
	"github.com/Andresch29/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerForProduct() *gin.Engine {
	db, _ := store.NewJsonStore("/Users/andhenao/Documents/bootcamp/goweb/products2.json")
	service := product.NewService(db)
	p := NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("", p.GetProducts())

	return r
}

func Test_GetAll(t *testing.T) {
	// Arrange
	server := createServerForProduct()

	request := httptest.NewRequest(http.MethodGet, "/products", nil)

	response := httptest.NewRecorder()
	
	// Act
	server.ServeHTTP(response, request)

	// Assert
	assert.Equal(t, http.StatusOK, response.Code)
}