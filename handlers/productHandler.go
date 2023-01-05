package handlers

import (
	"strconv"

	"github.com/Andresch29/go-web/db"
	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	products, err := db.GetAllProducts()
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

	product, ok, err := db.GetProductById(idProduct)

	if err != nil {
		ctx.String(500, "Error al leer la data")
		return
	}

	if !ok {
		ctx.String(404, "No se encontro el producto con id: %d", idProduct)
		return
	}

	ctx.JSON(200, product)
}

func GetProductByPrice(ctx *gin.Context) {
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.String(400, "El precio debe ser un numero")
		return
	}

	products, err := db.GetProductByPrice(price)
	if err != nil {
		ctx.JSON(500, nil)
		return
	}

	ctx.JSON(200, products)

}