package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"github.com/stripe/stripe-go/v78/webhook"
)

// PaymentRequest define a estrutura para criar um PaymentIntent
type PaymentRequest struct {
	Amount   int64  `json:"amount" binding:"required"`   // Valor em centavos (ex: 1000 = R$ 10,00)
	Currency string `json:"currency" binding:"required"` // Moeda (ex: "brl", "usd")
}

// CreatePaymentIntent cria um novo PaymentIntent no Stripe
func CreatePaymentIntent(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados invÃ¡lidos: " + err.Error()})
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(req.Currency),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// Cria o PaymentIntent no Stripe
	pi, err := paymentintent.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar intenÃ§Ã£o de pagamento: " + err.Error()})
		return
	}

	// Retorna o client_secret para o frontend completar o pagamento
	c.JSON(http.StatusOK, gin.H{
		"clientSecret": pi.ClientSecret,
		"id":           pi.ID,
	})
}

// HandleWebhook processa eventos enviados pelo Stripe (ex: confirmaÃ§Ã£o de pagamento)
func HandleWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler corpo da requisiÃ§Ã£o"})
		return
	}

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	// Verifica a assinatura do Stripe para garantir seguranÃ§a
	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assinatura do webhook invÃ¡lida"})
		return
	}

	// Processa o tipo de evento
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar JSON do Stripe"})
			return
		}
		// AQUI: Adicione a lÃ³gica para atualizar o banco de dados (ex: marcar pedido como pago)
		log.Printf("âœ… Pagamento %s recebido com sucesso!", paymentIntent.ID)

	case "payment_intent.payment_failed":
		log.Printf("â Œ Pagamento falhou.")

	default:
		log.Printf("â„¹ï¸  Evento nÃ£o tratado: %s", event.Type)
	}

	c.Status(http.StatusOK)
}
