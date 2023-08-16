package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"web/controller"
	"web/repository"
)

func main() {
	err := repository.InitTopicIndexMap("./data/")
	if err != nil {
		fmt.Println("topic初始化数据出错")
		os.Exit(-1)
	}

	if err := repository.InitPostIndexMap("./data/"); err != nil {
		fmt.Println("post初始化数据出错")
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

	// 回复话题（插入新帖）
	// 案例：POST http://localhost:9000/topic/1/post?content=youarewelcome
	r.POST("/topic/:id/post", func(c *gin.Context) {
		topicId := c.Param("id")
		content := c.Query("content")
		fmt.Println("/topic/id/post", topicId, content)

		data := controller.InsertNewPost(topicId, content)
		switch data.Code {
		case 0:
			c.JSON(200, data)
		case -1:
			c.JSON(http.StatusNotAcceptable, data)
		}
	})

	// r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务

	r.Run(":9000") // 指定在9000端口上 启动服务

	// http://localhost:9000
	// http://localhost:9000/ping
	// http://localhost:9000/topic/2
}
