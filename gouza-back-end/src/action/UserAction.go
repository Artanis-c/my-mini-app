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

func (ac *UserAction) Login(req req.WxLoginReq) result.Result {
	loginRes := ac.svc.WxLogin(req)

	return result.Success(login)
}
