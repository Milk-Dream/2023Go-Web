package main

import "github.com/gin-gonic/gin"


func main() {
	//gin-gonic.com
	//go get github.com/gin-gonic/gin

	/*
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"Hello World",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.String(200, "hello wrold")
	})

	//启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()

	*/
	//RESTFul api

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"POST",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"PUT",
		})
	})


	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"DELETE",
		})
	})

	//启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
	

}