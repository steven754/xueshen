package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取 c.Set("request", "中间件")
		status := c.Writer.Status()
		c.Next()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func A() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间价part1")
		c.Next()
		fmt.Println("中间件part2")

	}
}

func B() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间价part3")
		c.Next()
		fmt.Println("中间件part4")

	}
}
