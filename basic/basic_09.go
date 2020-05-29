package main

import (
	"github.com/gin-gonic/gin"
	gin_middlewares "go-abc/middle"
)
//中间件
//go不能动态对c.Context做方法的扩展。
//gin.Context的Set/Get可以用来交互接口响应数据给中间件。
func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), gin_middlewares.Default(map[string]interface{}{
		"code": 200,
		"msg":  "",
	}))
	router.GET("/", func(c *gin.Context) {
		c.Set("$.data.sub.sun.aaa.bbb.ccc.name", "abc")
	})
	router.Run(":1996")
}
