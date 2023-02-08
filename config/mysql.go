package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

//
//func Init() {
//	// 连接数据库
//	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Getconf()["User"], Getconf()["PassWord"], Getconf()["Host"], Getconf()["DataBase"])
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println("mysql连接失败，请检查")
//		log.Fatal(err)
//		return
//	}
//	fmt.Println("mysql连接成功")
//	DB = db
//	return
//}

// /初始化数据库链接
func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Getconf()["User"], Getconf()["PassWord"], Getconf()["Host"], Getconf()["DataBase"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			NoLowerCase:   false,
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Println("数据库连接异常", err)
		log.Fatal(err)
		return
	}
	fmt.Println("mysql连接成功")
	DB = db
	return
}
