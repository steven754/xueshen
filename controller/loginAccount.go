package controller

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"xueshen/config"
//	"xueshen/model"
//)
//
//func LoginAccount(c *gin.Context) {
//	// 1. 从请求中把数据拿出来
//	var loginaccount := model.CreateSignAccount()
//
//	err := c.BindJSON(&loginaccount)
//	if err != nil {
//		fmt.Println("bindJson错误", err)
//		return
//	}
//	if loginaccount.Account == "" || loginaccount.Password == "" {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code": 1001,
//			"msg":  "账号密码不能为空",
//		})
//	}
//
//
//
//
//
//
//
//	else if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1000,
//			"msg":  "fail",
//			"info": err.Error()})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"msg":  "success",
//			"data": loginaccount,
//		})
//	}
//}
