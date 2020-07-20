package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pgmot/awesomelib"
	"github.com/pgmot/superwebapp/pkg/hello"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": hello.Hello(name),
		})
	})

	r.GET("/data", func(c *gin.Context){
		data := awesomelib.NewHoge("mot")

		c.JSON(http.StatusOK, data)
	})

	panic(r.Run())
}