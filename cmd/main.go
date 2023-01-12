package main

import (
	"log"

	"github.com/Andresch29/go-web/cmd/handlers"
	"github.com/Andresch29/go-web/internal/product"
	"github.com/Andresch29/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al caergar el archivo .env")
	} 

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	repo, err := store.NewJsonStore("/Users/andhenao/Documents/bootcamp/goweb/products2.json")
	if err != nil {
		log.Fatal("Error al conectar la db")
	}
	
	service := product.NewService(repo)
	product := handlers.NewProduct(service)

	productRouter := server.Group("/products")
	productRouter.GET("", product.GetProducts)
	productRouter.GET("/:id", product.GetProductById)
	productRouter.GET("/search", product.GetProductByPrice)
	productRouter.POST("", product.Create)
	productRouter.PUT("/:id", product.Update)
	productRouter.DELETE("/:id", product.Delete)


	server.Run()
}
