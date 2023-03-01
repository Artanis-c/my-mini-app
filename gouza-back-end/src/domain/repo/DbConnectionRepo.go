package repo

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_CONN *gorm.DB

func NewDbConnection() *gorm.DB {
	if DB_CONN == nil {
		dsn := "root:Admin123*@tcp(127.0.0.1:3306)/gouza?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		DB_CONN = db
		return DB_CONN
	} else {
		return DB_CONN
	}
}
