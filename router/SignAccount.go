package router

import (
	"github.com/gin-gonic/gin"
	"xueshen/controller"
)

func CreateSignAccount(r *gin.Engine) *gin.Engine {
	r.POST("/signup", controller.CreateSignAccount)
	return r
}
