package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//React build path
	router.Use(static.Serve("/", static.LocalFile("./app/build", true)))

	router.GET("/date", post)
	router.POST("/message", get)

	router.Run(":3000")
}

func post(c *gin.Context) {
	input := c.PostForm("input")
	c.String(http.StatusOK, "map array: \n%s", input)
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"date": time.Now().Format("2006-01-02"),
	})
}
