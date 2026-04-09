package main

import (
	"api-focus/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	gin.SetMode(gin.DebugMode) // ou ReleaseMode depois

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Focus rodando?",
		})
	})

	r.Run(":8080")
}