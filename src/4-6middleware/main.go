package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 使用中间件
	r.Use(func(c *gin.Context) {
		fmt.Println("middleware start", c.Request.URL)
		c.Next()
		fmt.Println("middleware over")
	})

	r.POST("/login", func(c *gin.Context) {
		fmt.Println("在login处理函数")
		c.Data(200, "text/plain; charset=utf-8", []byte("login"))
	})

	r.POST("/logout", func(c *gin.Context) {
		fmt.Println("在logout处理函数")
		c.Data(200, "text/plain; charset=utf-8", []byte("logout"))
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
