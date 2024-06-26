package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET(
		"/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hey Go URL Shortner !",
			})
		})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web server - Error: %v", err))
	}
}
