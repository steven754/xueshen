package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"xueshen/config"
	"xueshen/middle"
	"xueshen/model"
	"xueshen/router"
)

func main() {
	config.Init()
	r := gin.Default()
	r.Use(middle.Auth())
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("文件关闭错误", err)
		}
	}(f)
	err := config.DB.AutoMigrate(&model.SignAccount{})
	if err != nil {
		fmt.Println("自动建表错误", err)
		return
	}
	// 注册路由
	r = router.CreateSignAccount()
	r = router.LoginAccount()
	fmt.Println(config.Strval(config.Getconf()["Url"]))
	err = r.Run(config.Strval(config.Getconf()["Url"]))
	if err != nil {
		fmt.Println("启动失败", err)
		return
	}
}
