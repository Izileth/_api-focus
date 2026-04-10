package handlers

import (
	"api-focus/internal/database"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns the status of the API and database
func HealthCheck(c *gin.Context) {
	dbStatus := "UP"
	if database.DB == nil {
		dbStatus = "DOWN"
	} else {
		err := database.DB.Ping(context.Background())
		if err != nil {
			dbStatus = "DOWN"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"api":      "UP",
		"database": dbStatus,
		"message":  "API Focus systems check",
	})
}

// RootInfo returns information about the API
func RootInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to API Focus",
		"versions": []string{"v1", "v2", "v3"},
	})
}
