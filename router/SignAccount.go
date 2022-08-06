package router

import (
	"github.com/gin-gonic/gin"
	"xueshen/controller"
)

func CreateSignAccount() *gin.Engine {
	r := gin.Default()
	r.POST("/signup", controller.CreateSignAccount)
	return r
}
