package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Getconf()["User"], Getconf()["PassWord"], Getconf()["Host"], Getconf()["DataBase"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql连接失败，请检查")
		log.Fatal(err)
		return
	}
	fmt.Println("mysql连接成功")
	DB = db
	return
}
