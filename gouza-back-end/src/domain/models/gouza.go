package models

import "time"

// User 用户表
type User struct {
	ID         uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"id"`
	OpenID     string    `gorm:"column:open_id;type:varchar(70);not null" json:"openId"`
	NickName   string    `gorm:"column:nick_name;type:varchar(30)" json:"nickName"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID         string
	OpenID     string
	NickName   string
	CreateTime string
}{
	ID:         "id",
	OpenID:     "open_id",
	NickName:   "nick_name",
	CreateTime: "create_time",
}
