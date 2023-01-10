package main

import (
	"github.com/Andresch29/go-web/cmd/handlers"
	"github.com/Andresch29/go-web/internal/product"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	repo := product.NewMemory()
	service := product.NewService(repo)
	product := handlers.NewProduct(service)

	productRouter := server.Group("/products")
	productRouter.GET("", product.GetProducts)
	productRouter.GET("/:id", product.GetProductById)
	productRouter.GET("/search", product.GetProductByPrice)
	productRouter.POST("", product.Create)


	server.Run()
}
