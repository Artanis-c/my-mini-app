package router

import (
	"github.com/gin-gonic/gin"
	"gouza-back-end/src/domain/ioc"
	req2 "gouza-back-end/src/domain/models/req"
	"net/http"
)

const CREATE_USER_ROUTE string = "/api/user/create"

func RegisterRoute(webServer *gin.Engine) {
	webServer.POST(CREATE_USER_ROUTE, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		var req req2.CreateUserReq
		context.BindJSON(&req)
		context.JSON(http.StatusOK, action.CrateUserInfo(req))
	})
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "" {
			context.JSON(http.StatusUnauthorized, nil)
		} else {
			context.Next()
		}
	}
}
