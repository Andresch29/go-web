package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/Andresch29/go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != requiredToken {
			web.NewNotAuthResponse(ctx, http.StatusUnauthorized, "No estas autorizado")
			return 
		}

		ctx.Next()
	}
}

func ServerLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("[GIN] Method: %s - URL: localhost%s - LEN: %d", ctx.Request.Method,ctx.Request.RequestURI, ctx.Request.ContentLength)

		ctx.Next()
	}
}