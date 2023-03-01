package repo

import (
	"gorm.io/gorm"
	"gouza-back-end/src/domain/models"
	"gouza-back-end/src/domain/models/req"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUserInfo(req req.CreateUserReq) models.User {
	user := models.User{OpenID: req.OpenId, CreateTime: time.Now()}
	repo.db.Create(&user)
	return user
}
