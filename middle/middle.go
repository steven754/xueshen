package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件
func LoginMiddle(ctx *gin.Context) {
	login := ctx.Query("login")
	if login != "" {
		ctx.Next()
		ctx.String(http.StatusOK, "Hi, "+login+",欢迎访客登录")
	} else {
		ctx.Abort()
		ctx.String(http.StatusForbidden, "未登录，拒绝访问")
	}
}
