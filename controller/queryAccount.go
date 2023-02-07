package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xueshen/model"
)

func QueryAccount(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var accountlist model.AccountList
	err := c.BindJSON(&accountlist)
	if err != nil {
		fmt.Println("bindJson错误", err)
		return
	}
	model.QueryAccountList(&accountlist)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  1000,
			"msg":   "fail",
			"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": accountlist,
		})
	}
}
