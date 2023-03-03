package repo

import (
	"gorm.io/gorm"
	"gouza-back-end/src/domain/common"
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

func (repo *UserRepository) CreateUserInfo(req req.CreateUserReq) *models.User {
	var now = time.Now()
	user := &models.User{OpenID: req.OpenId, CreateTime: &now}
	repo.db.Create(user)
	return user
}

func (repo *UserRepository) QueryUserInfo(openId string) *models.User {
	user := models.User{}
	repo.db.Where("open_id=?", openId).Find(&user)
	return &user
}

func (repo *UserRepository) QueryUserInfoExist(openId string) bool {
	var count int64
	repo.db.Where("open_id=?", openId).Count(&count)
	return count > 0
}

func (repo *UserRepository) QueryWxInfo() *models.WeChaetInfo {
	wxInfo := &models.WeChaetInfo{}
	if common.APP_ID == "" {
		repo.db.Where(models.WeChaetInfo{}).First(wxInfo)
		common.APP_ID = wxInfo.AppID
		common.APP_SECRET = wxInfo.AppSecret
		return wxInfo
	} else {
		wxInfo.AppID = common.APP_ID
		wxInfo.AppSecret = common.APP_SECRET
		return wxInfo
	}
}
