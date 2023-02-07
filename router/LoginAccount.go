package router

import (
	"github.com/gin-gonic/gin"
	"xueshen/middle"
)

func LoginAccount(r *gin.Engine) *gin.Engine {
	r.Use(middle.Auth())

	return r
}

func UpdateAccountPassword(r *gin.Engine) *gin.Engine {
	r.Use(middle.Auth())

	return r
}
