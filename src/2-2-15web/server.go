package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"web/controller"
	"web/repository"
)

func main() {
	err := repository.InitTopicIndexMap("./data/")
	if err != nil {
		fmt.Println("初始化数据出错")
		os.Exit(-1)
	}

	r := gin.Default()
	// url - ping
	r.GET("/ping", func(c *gin.Context) {
		// - 200 是请求成功状态码，问就是规定200是成功，404是page not found
		// - type gin.H map[string]any
		// c.JSON 快速使用map或者struct返回json数据
		// c.JSON serializes the given struct as JSON into the response body.
		// It also sets the Content-Type as "application/json".
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// url - /topic/id参数
	r.GET("/topic/:id", func(c *gin.Context) {
		// 获取id参数
		topicId := c.Param("id")
		// 通过controller层的QueryPageInfo函数找Topic相关数据
		data := controller.QueryPageInfo(topicId)
		// 直接把数据返回
		c.JSON(200, data)
	})
	// r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务

	r.Run(":9000") // 指定在9000端口上 启动服务

	// http://localhost:9000
	// http://localhost:9000/ping
	// http://localhost:9000/topic/2
}
