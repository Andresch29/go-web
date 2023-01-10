package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Andresch29/go-web/internal/domain"
	"github.com/Andresch29/go-web/internal/product"
	"github.com/gin-gonic/gin"
)

type productRequest struct {
	Name 		string `json:"name" binding:"required"`
	Quantity 	int `json:"quantity" binding:"required"`
	CodeValue 	string `json:"code_value" binding:"required"`
	IsPublished bool `json:"is_published"`
	Expiration 	string `json:"expiration" binding:"required"`
	Price 		float64 `json:"price" binding:"required"`
}


type Product struct {
	service product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{service}
}

func (p *Product) GetProducts(ctx *gin.Context) {
	products := p.service.GetAll()
	
	ctx.JSON(200, products)
}

func (p *Product) GetProductById(ctx *gin.Context) {
	idProduct, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(400, "El id debe ser un numero")
		return
	}

	product, ok := p.service.GetById(idProduct)

	if !ok {
		ctx.String(404, "No se encontro el producto con id: %d", idProduct)
		return
	}

	ctx.JSON(200, product)
}

func (p *Product) GetProductByPrice(ctx *gin.Context) {
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.String(400, "El precio debe ser un numero")
		return
	}

	products:= p.service.GetByPrice(price)

	ctx.JSON(200, products)

}

func (p *Product) Create(ctx *gin.Context) {
	var productRequest productRequest

	err := ctx.ShouldBindJSON(&productRequest)
	if err != nil {
		ctx.String(400, "Error en la peticion")
		return
	}

	if _, err := time.Parse("02/01/2006", productRequest.Expiration); err != nil {
		ctx.String(400, "Fecha invalida")
		return
	}

	product := &domain.Product{
		Name: productRequest.Name,
		Quantity: productRequest.Quantity,
		CodeValue: productRequest.CodeValue,
		IsPublished: productRequest.IsPublished,
		Expiration: productRequest.Expiration,
		Price: productRequest.Price,
	}

	productDB, err := p.service.Create(product)
	if err != nil {
		ctx.String(500, "Error al crear producto")
		return
	}

	ctx.JSON(http.StatusCreated, productDB)
}