package handlers

import (
	"api-focus/internal/stripe"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/webhook"
)

type PaymentRequest struct {
	Amount   int64  `json:"amount" binding:"required"`
	Currency string `json:"currency" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type CheckoutRequest struct {
	Amount     int64  `json:"amount" binding:"required"`
	Currency   string `json:"currency" binding:"required"`
	SuccessURL string `json:"success_url" binding:"required"`
	CancelURL  string `json:"cancel_url" binding:"required"`
}

// ✅ PaymentIntent moderno
func CreatePaymentIntent(c *gin.Context) {
	var req PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pi, err := stripeclient.CreatePaymentIntent(req.Amount, req.Currency, req.Email)
	if err != nil {
		// Tratamento correto de erro do Stripe
		if stripeErr, ok := err.(*stripe.Error); ok {
			log.Printf("❌ Erro no Stripe (CreatePaymentIntent): Code=%s, DeclineCode=%s, Msg=%s", 
				stripeErr.Code, stripeErr.DeclineCode, stripeErr.Msg)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": stripeErr.Msg,
			})
			return
		}

		log.Printf("❌ Erro interno (CreatePaymentIntent): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"clientSecret": pi.ClientSecret,
		"id":           pi.ID,
		"status":       pi.Status,
	}

	if pi.NextAction != nil {
		response["nextAction"] = pi.NextAction
	}

	c.JSON(http.StatusOK, response)
}

// ✅ Checkout Session
func CreateCheckoutSession(c *gin.Context) {
	var req CheckoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := stripeclient.CreateCheckoutSession(
		req.Amount,
		req.Currency,
		req.SuccessURL,
		req.CancelURL,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": session.URL,
		"id":  session.ID,
	})
}

// ✅ Webhook
func HandleWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler corpo"})
		return
	}

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Webhook inválido"})
		return
	}

	switch event.Type {

	case "payment_intent.succeeded":
		var pi stripe.PaymentIntent

		if err := json.Unmarshal(event.Data.Raw, &pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro JSON"})
			return
		}

		log.Printf("✅ Pagamento %s confirmado", pi.ID)

	case "payment_intent.payment_failed":
		var pi stripe.PaymentIntent
		if err := json.Unmarshal(event.Data.Raw, &pi); err == nil {
			log.Printf("❌ Pagamento %s falhou: Code=%s, Message=%s", 
				pi.ID, pi.LastPaymentError.Code, pi.LastPaymentError.Message)
		} else {
			log.Println("❌ Pagamento falhou (erro ao ler detalhes)")
		}

	default:
		log.Println("ℹ️ Evento ignorado:", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}