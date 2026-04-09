package main

import (
	"api-focus/internal/config"
	"api-focus/internal/database"
	stripeclient "api-focus/internal/stripe"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("âš™ï¸  Iniciando a API...")

	// Inicializa configuraÃ§Ãµes de ambiente
	config.LoadEnv()

	// Inicializa conexÃµes e validaÃ§Ãµes
	database.Connect()
	stripeclient.Init()

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Focus rodando!",
		})
	})

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("âš¡ï¸  API pronta e ouvindo na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("â Œ Falha ao iniciar o servidor: %v", err)
	}
}