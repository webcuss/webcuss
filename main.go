package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/sup", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"token": "",
		})
	})

	r.POST("/sin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": "",
		})
	})

	r.POST("/sout", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
