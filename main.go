package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pgmot/superwebapp/pkg/hello"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": hello.Hello(name),
		})
	})

	panic(r.Run())
}