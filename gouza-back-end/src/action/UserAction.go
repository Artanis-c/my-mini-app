package action

import (
	result "gouza-back-end/src/domain/models/dto"
	"gouza-back-end/src/domain/models/req"
	"gouza-back-end/src/domain/service"
)

type UserAction struct {
	svc *service.UserService
}

func NewUserAction(service *service.UserService) *UserAction {
	return &UserAction{
		svc: service,
	}
}

func (ac *UserAction) CrateUserInfo(req req.CreateUserReq) result.Result {
	user := ac.svc.CreateUser(req)
	return result.Success(user)
}

// Login 用户登录
func (ac *UserAction) Login(req req.WxLoginReq) result.Result {
	loginRes, err := ac.svc.WxLogin(req)
	if err != nil {
		return result.Fail(err.Error())
	}
	return result.Success(loginRes)
}
