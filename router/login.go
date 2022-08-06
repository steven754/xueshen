package router

import (
	"github.com/gin-gonic/gin"
	"xueshen/controller"
)

func CreateLoginAccount() *gin.Engine {
	r := gin.Default()
	r.POST("/signup", controller.CreateLoginAccount)
	return r
}
