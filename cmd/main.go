package main

import (
	"log"
	"os"

	"github.com/Andresch29/go-web/cmd/handlers"
	"github.com/Andresch29/go-web/cmd/middleware"
	"github.com/Andresch29/go-web/docs"
	"github.com/Andresch29/go-web/internal/product"
	"github.com/Andresch29/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This Api Handle MELI Products.
// @termsOfService MELI

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al caergar el archivo .env")
	} 

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.ServerLog())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
	productRouter.GET("", product.GetProducts())
	productRouter.GET("/:id", product.GetProductById())
	productRouter.GET("/search", product.GetProductByPrice())
	productRouter.Use(middleware.TokenAuthMiddleware())
	productRouter.POST("", product.Create())
	productRouter.PUT("/:id", product.Update())
	productRouter.DELETE("/:id", product.Delete())


	server.Run()
}
