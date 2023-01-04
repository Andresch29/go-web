package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Andresch29/go-web/handlers"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	productRouter := server.Group("/products")
	productRouter.GET("", handlers.GetProducts)
	productRouter.GET("/:id", handlers.GetProductById)
	productRouter.GET("/search", handlers.GetProductByPrice)



	server.Run()
}
