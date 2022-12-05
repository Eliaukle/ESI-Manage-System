package common

import (
	"blog_server/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	user := "root"
	password := "root password"
	host := "localhost"
	port := "3306"
	database := "esi_db"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		database,
		charset)
	// 连接数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	// 迁移数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
