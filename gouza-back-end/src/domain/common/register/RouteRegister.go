package register

import (
	"github.com/gin-gonic/gin"
	"gouza-back-end/src/domain/common/router"
	"gouza-back-end/src/domain/ioc"
	"gouza-back-end/src/domain/models/req"
	"net/http"
)

func RegisterRoute(webServer *gin.Engine) {
	webServer.POST(router.CREATE_USER_ROUTE, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		var req req.CreateUserReq
		context.BindJSON(&req)
		context.JSON(http.StatusOK, action.CrateUserInfo(req))
	})
	webServer.POST(router.WX_LOGIN, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		var req req.WxLoginReq
		context.BindJSON(&req)
		context.JSON(http.StatusOK, action.Login(req))
	})
}
