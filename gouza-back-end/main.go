package main

import (
	"github.com/gin-gonic/gin"
	"gouza-back-end/src/domain/common/logutils"
	"gouza-back-end/src/domain/common/router"
	"gouza-back-end/src/domain/ioc"
	"gouza-back-end/src/domain/models/req"
	"net/http"
)

func main() {
	webServer := gin.New()
	webServer.Use(gin.Recovery())
	webServer.Use(router.AuthMiddleWare())
	webServer.POST(router.CREATE_USER_ROUTE, func(context *gin.Context) {
		var createUserReq req.CreateUserReq
		err := context.BindJSON(&createUserReq)
		if err != nil {
			logutils.LogInfo("创建用户绑定失败")
		}
		action, err := ioc.BuildUserAction()
		info := action.CrateUserInfo(createUserReq)
		context.JSON(http.StatusOK, &info)
	})
	webServer.Run(":8080")
}
