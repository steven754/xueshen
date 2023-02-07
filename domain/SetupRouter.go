package domain

import (
	"github.com/gin-gonic/gin"
	"xueshen/controller"
	"xueshen/middle"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middle.Auth())
	r.POST("/login", controller.LoginAccount)
	r.POST("/update/pswd", controller.UpdateAccountPassword)
	r.POST("/sign", controller.CreateSignAccount)
	r.POST("/query", controller.QueryAccount)
	return r
}
