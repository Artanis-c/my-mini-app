package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const CREATE_USER_ROUTE string = "/api/user/create"

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "" {
			context.Abort()
			context.JSON(http.StatusUnauthorized, nil)
		} else {
			context.Next()
		}
	}
}
