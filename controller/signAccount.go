package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xueshen/model"
)

func CreateSignAccount(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var signaccount model.SignAccount
	err := c.BindJSON(&signaccount)
	if err != nil {
		fmt.Println("bindJson错误", err)
		return
	}
	if signaccount.Account == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1001,
			"msg":  "账号不能为空",
		})
	} else {
		// 2. 存入数据库
		err := model.CreateSignAccount(&signaccount)
		if strings.Contains(err.Error(), "Error 1062") {
			c.JSON(http.StatusOK, gin.H{
				"code": 1000,
				"msg":  "该账号已注册",
			})
		} else if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  1000,
				"msg":   "fail",
				"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": signaccount,
			})
		}
	}
}
