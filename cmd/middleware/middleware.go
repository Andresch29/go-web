package middleware

import (
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