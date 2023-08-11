package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		//	前置处理
		fmt.Println("login request")

		c.Data(200, "text/plain; charset=utf-8", []byte("login"))

		//	后置处理
		fmt.Println("login sucess")
	})

	r.POST("/logout", func(c *gin.Context) {
		//	前置处理
		fmt.Println("logout request")

		c.Data(200, "text/plain; charset=utf-8", []byte("logout"))

		//	后置处理
		fmt.Println("logout success")
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
