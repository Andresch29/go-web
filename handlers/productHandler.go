package handlers

import (
	"strconv"

	"github.com/Andresch29/go-web/db"
	"github.com/Andresch29/go-web/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	products, err := db.ReadFile("/Users/andhenao/Documents/bootcamp/goweb/products.json")
	if err != nil {
		ctx.JSON(500, nil)
		return
	}
	
	ctx.JSON(200, products)
}

func GetProductById(ctx *gin.Context) {
	idProduct, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(400, "El id debe ser un numero")
		return
	}

	products, err := db.ReadFile("/Users/andhenao/Documents/bootcamp/goweb/products.json")
	if err != nil {
		ctx.String(500, "Error al leer la data")
		return
	}

	var product = models.Product{}
	ok := false
	for _, p := range *products {
		if p.Id == idProduct {
			product = p
			ok = true
			break
		}
	}
	
	if !ok {
		ctx.String(404, "No se encontro el producto con id: %d", idProduct)
		return
	}

	ctx.JSON(200, product)
}

func GetProductByPrice(ctx *gin.Context) {
	products, err := db.ReadFile("/Users/andhenao/Documents/bootcamp/goweb/products.json")
	if err != nil {
		ctx.JSON(500, nil)
		return
	}

	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.String(400, "El precio debe ser un numero")
		return
	}

	productsResponse := models.Products{}
	for _, p := range *products {
		if p.Price > price {
			productsResponse = append(productsResponse, p)
		}
	}

	ctx.JSON(200, productsResponse)

}