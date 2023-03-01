package service

import (
	"gouza-back-end/src/domain/models"
	"gouza-back-end/src/domain/models/req"
	"gouza-back-end/src/domain/repo"
)

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

//CreateUser 创建用户
func (s *UserService) CreateUser(req req.CreateUserReq) models.User {
	return s.userRepo.CreateUserInfo(req)
}
