package router

import (
	"github.com/gin-gonic/gin"
	"xueshen/controller"
)

func LoginAccount(r *gin.Engine) *gin.Engine {
	r.POST("/login", controller.LoginAccount)
	return r
}

func UpdateAccountPassword(r *gin.Engine) *gin.Engine {
	r.POST("/update/accountpwd", controller.UpdateAccountPassword)
	return r
}
