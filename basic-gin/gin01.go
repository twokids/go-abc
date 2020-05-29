package main

import "github.com/gin-gonic/gin"

type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c)
	}
}

func user(c *gin.Context) error {
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

	r.GET("/ping", wrapper(user))

	r.Run() // 在 0.0.0.0:8080 上监听并服务
}
