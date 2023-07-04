package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controllers struct {
	Router *gin.Engine
}

func New() *Controllers {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/test", func(context *gin.Context) {
			context.Status(http.StatusOK)
		})
	}

	return &Controllers{Router: r}
}
