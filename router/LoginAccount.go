package router

import (
	"github.com/gin-gonic/gin"
	"xueshen/controller"
)

func LoginAccount() *gin.Engine {
	r := gin.Default()
	r.POST("/login", controller.LoginAccount)
	return r
}

func UpdateAccountPassword() *gin.Engine {
	r := gin.Default()
	r.POST("/update/accountpwd", controller.UpdateAccountPassword)
	return r
}
