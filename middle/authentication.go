package middle

import (
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) (username interface{}, ok bool) {
	//获取cookie
	value, _ := c.Cookie("username")
	//解码
	store, _ := base64.StdEncoding.DecodeString(value)
	//session验证
	session := sessions.Default(c)
	username = session.Get("username")
	return username, username == string(store)
}
