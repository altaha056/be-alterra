package config

import (
	"project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initMigrate() {
	DB.AutoMigrate(&model.Users{})
}

func InitDB() {
	dsn := "root@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}