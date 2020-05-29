package main

import "github.com/gin-gonic/gin"

type HandlerFunc1 func(c *gin.Context) error

func wrapper1(handler HandlerFunc1) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c)
	}
}

func user1(c *gin.Context) error {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return nil
}

func main() {

	r := gin.Default()
	r.GET("/go", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "go",
		})
	})

	r.GET("/ping", wrapper1(user1))

	r.Run() // 在 0.0.0.0:8080 上监听并服务
}
