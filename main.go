package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lylx0/gin-demo/data"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/products", data.GetProducts)
	r.Run()
}
