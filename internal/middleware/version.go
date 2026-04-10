package middleware

import (
	"github.com/gin-gonic/gin"
)

// VersionMiddleware adds the API version to the response header
func VersionMiddleware(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-API-Version", version)
		c.Next()
	}
}
