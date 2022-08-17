package router

import (
	"github.com/gin-gonic/gin"
	"xueshen/controller"
	"xueshen/middle"
)

func LoginAccount(r *gin.Engine) *gin.Engine {
	r.Use(middle.Auth())
	r.POST("/login", controller.LoginAccount)
	return r
}

func UpdateAccountPassword(r *gin.Engine) *gin.Engine {
	r.Use(middle.Auth())
	r.POST("/update/pswd", controller.UpdateAccountPassword)
	return r
}
