package web

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, status int, err string) {
	reponseJson := errorResponse{
		Status: status,
		Message: err,
	}

	ctx.JSON(status, reponseJson)
}

type response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context, status int, data interface{}) {
	reponseJson := response{
		Status: status,
		Data: data,
	}

	ctx.JSON(status, reponseJson)
}