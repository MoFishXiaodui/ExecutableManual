package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/sis", func(c *gin.Context) {
		c.Data(200, "text/plain; charset=utf-8", []byte("ok"))
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
