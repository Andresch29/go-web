package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Andresch29/go-web/internal/domain"
	"github.com/Andresch29/go-web/internal/product"
	"github.com/Andresch29/go-web/pkg/web"
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

// @summary List products
// @tags Products
// @Description get products
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Router /products [get]
func (p *Product) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll()
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		web.NewResponse(ctx, http.StatusOK, products)
	}
}

func (p *Product) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idProduct, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "El id debe ser un entero")
			return
		}

		product, err := p.service.GetById(idProduct)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error en el servidor")
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		if product == nil {
			ctx.String(http.StatusNotFound, "No se encontro el producto")
			web.NewErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("No existe el producto con %d", idProduct))
			return
		}

		web.NewResponse(ctx, http.StatusOK, product)
	}
}

func (p *Product) GetProductByPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "El precio debe ser un valor numerico")
			return
		}

		products, err := p.service.GetByPrice(price)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		web.NewResponse(ctx, http.StatusOK, products)
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productRequest productRequest

		err := ctx.ShouldBindJSON(&productRequest)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "Error en la peticion")
			return
		}

		if _, err := time.Parse("02/01/2006", productRequest.Expiration); err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "Error en la peticion")
			return
		}

		existsCode, err := p.service.ExistsByCode(productRequest.CodeValue)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error en el servidor")
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		if existsCode {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "Error en la peticion")
			return
		}

		product := domain.Product{
			Name: productRequest.Name,
			Quantity: productRequest.Quantity,
			CodeValue: productRequest.CodeValue,
			IsPublished: productRequest.IsPublished,
			Expiration: productRequest.Expiration,
			Price: productRequest.Price,
		}

		productDB, err := p.service.Create(product)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		web.NewResponse(ctx, http.StatusCreated, productDB)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "El id debe ser un numero")
			return
		}

		var productRequest domain.Product
		err = ctx.ShouldBindJSON(&productRequest)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "Error en la peticion")
			return
		}

		productNew := domain.Product{
			Id: id,
			Name: productRequest.Name,
			Quantity: productRequest.Quantity,
			CodeValue: productRequest.CodeValue,
			IsPublished: productRequest.IsPublished,
			Expiration: productRequest.Expiration,
			Price: productRequest.Price,
		}

		productResponse, err := p.service.Update(productNew)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		if productResponse == nil {
			web.NewErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("No existe el producto con %d", id))
			return
		}

		web.NewResponse(ctx, http.StatusOK, productResponse)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		idProduct, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusBadRequest, "El id debe ser un numero")
			return
		}

		productDB, err := p.service.GetById(idProduct)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}
		if productDB == nil {
			web.NewErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("No existe el producto con %d", idProduct))
			return
		}

		deleted, err := p.service.Delete(idProduct)
		if err != nil {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		if !deleted {
			web.NewErrorResponse(ctx, http.StatusInternalServerError, "Error en el servidor")
			return
		}

		web.NewResponse(ctx, http.StatusNoContent, "Producto eliminado")
	}
}
