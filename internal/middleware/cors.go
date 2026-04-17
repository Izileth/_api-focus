package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles Cross-Origin Resource Sharing

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		allowedOrigins := map[string]bool{
			"http://localhost:1222": true,
			"https://modusfocus.online": true,
			"http://modusfocus.online": true,
		}

		// 🔥 SEMPRE responde OPTIONS com headers
		if c.Request.Method == "OPTIONS" {
			if allowedOrigins[origin] {
				c.Header("Access-Control-Allow-Origin", origin)
			} else {
				c.Header("Access-Control-Allow-Origin", "*") // fallback p/ não quebrar
			}

			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

			c.AbortWithStatus(204)
			return
		}

		// Requisições normais
		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		c.Header("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}