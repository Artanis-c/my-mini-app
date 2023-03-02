package action

import (
	"github.com/gin-gonic/gin"
	"gouza-back-end/src/domain/common/router"
	"gouza-back-end/src/domain/ioc"
	result "gouza-back-end/src/domain/models/dto"
	"gouza-back-end/src/domain/models/req"
	"gouza-back-end/src/domain/service"
	"net/http"
)

type UserAction struct {
	svc *service.UserService
}

func NewUserAction(service *service.UserService) *UserAction {
	return &UserAction{
		svc: service,
	}
}

func RegisterRoute(webServer *gin.Engine) {
	webServer.POST(router.CREATE_USER_ROUTE, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		var req req.CreateUserReq
		context.BindJSON(&req)
		context.JSON(http.StatusOK, action.CrateUserInfo(req))
	})
}

func (ac *UserAction) CrateUserInfo(req req.CreateUserReq) result.Result {
	user := ac.svc.CreateUser(req)
	return result.Success(user)
}
