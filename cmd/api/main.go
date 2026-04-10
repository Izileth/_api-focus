package main

import (
	"api-focus/internal/config"
	"api-focus/internal/database"
	"api-focus/internal/handlers"
	"api-focus/internal/middleware"
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

	// Endpoints globais
	r.GET("/", handlers.RootInfo)
	r.GET("/health", handlers.HealthCheck)

	// API Versioning
	api := r.Group("/api")
	{
		// Version 1
		v1 := api.Group("/v1")
		v1.Use(middleware.VersionMiddleware("v1"))
		{
			v1.GET("/health", handlers.HealthCheck)
			payments := v1.Group("/payments")
			{
				payments.POST("/create-intent", handlers.CreatePaymentIntent)
				payments.POST("/webhook", handlers.HandleWebhook)
			}
		}

		// Version 2
		v2 := api.Group("/v2")
		v2.Use(middleware.VersionMiddleware("v2"))
		{
			v2.GET("/health", handlers.HealthCheck)
			payments := v2.Group("/payments")
			{
				payments.POST("/create-intent", handlers.CreatePaymentIntent)
				payments.POST("/webhook", handlers.HandleWebhook)
			}
		}

		// Version 3
		v3 := api.Group("/v3")
		v3.Use(middleware.VersionMiddleware("v3"))
		{
			v3.GET("/health", handlers.HealthCheck)
			payments := v3.Group("/payments")
			{
				payments.POST("/create-intent", handlers.CreatePaymentIntent)
				payments.POST("/webhook", handlers.HandleWebhook)
			}
		}
	}

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("âš¡ï¸  API pronta e ouvindo na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("â Œ Falha ao iniciar o servidor: %v", err)
	}
}