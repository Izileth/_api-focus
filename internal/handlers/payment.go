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

// PaymentRequest define a estrutura para criar um PaymentIntent
type PaymentRequest struct {
	Amount        int64  `json:"amount" binding:"required"`         // Valor em centavos (ex: 1000 = R$ 10,00)
	Currency      string `json:"currency" binding:"required"`       // Moeda (ex: "brl", "usd")
	PaymentMethod string `json:"payment_method" binding:"required"` // "card", "pix", "boleto"
	Email         string `json:"email"`                             // Necessário para Pix e Boleto
	Name          string `json:"name"`                              // Necessário para Pix e Boleto
	TaxID         string `json:"tax_id"`                            // CPF/CNPJ (necessário para Boleto e Pix)
}

// CheckoutRequest define a estrutura para criar uma Checkout Session
type CheckoutRequest struct {
	Amount     int64  `json:"amount" binding:"required"`
	Currency   string `json:"currency" binding:"required"`
	SuccessURL string `json:"success_url" binding:"required"`
	CancelURL  string `json:"cancel_url" binding:"required"`
}

// CreateCheckoutSession cria uma sessão de checkout e retorna a URL
func CreateCheckoutSession(c *gin.Context) {
	var req CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	session, err := stripeclient.CreateCheckoutSession(req.Amount, req.Currency, req.SuccessURL, req.CancelURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar sessão de checkout: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": session.URL,
		"id":  session.ID,
	})
}

// CreatePaymentIntent cria um novo PaymentIntent no Stripe usando o stripeclient
func CreatePaymentIntent(c *gin.Context) {
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados invÃ¡lidos: " + err.Error()})
		return
	}

	var pi *stripe.PaymentIntent
	var err error

	switch req.PaymentMethod {
	case "pix":
		if req.Email == "" || req.Name == "" || req.TaxID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email, Nome e TaxID (CPF/CNPJ) sÃ£o obrigatÃ³rios para Pix"})
			return
		}
		if req.Currency != "brl" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Pix suporta apenas a moeda BRL"})
			return
		}
		pi, err = stripeclient.CreatePixIntent(req.Amount, req.Currency, req.Email, req.Name, req.TaxID)

	case "boleto":
		if req.Email == "" || req.Name == "" || req.TaxID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email, Nome e TaxID (CPF/CNPJ) sÃ£o obrigatÃ³rios para Boleto"})
			return
		}
		pi, err = stripeclient.CreateBoletoIntent(req.Amount, req.Currency, req.Email, req.Name, req.TaxID)

	case "card":
		pi, err = stripeclient.CreateCardIntent(req.Amount, req.Currency)

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "MÃ©todo de pagamento nÃ£o suportado: " + req.PaymentMethod})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar intenÃ§Ã£o de pagamento: " + err.Error()})
		return
	}

	// Retorna o client_secret e outras informaÃ§Ãµes necessÃ¡rias (como QR Code para Pix)
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
		log.Printf("âœ… Pagamento %s de %d %s recebido com sucesso!", paymentIntent.ID, paymentIntent.Amount, paymentIntent.Currency)
		
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Pagamento confirmado com sucesso",
			"id":      paymentIntent.ID,
		})
		return

	case "payment_intent.payment_failed":
		log.Printf("â Œ Pagamento falhou.")
		c.JSON(http.StatusOK, gin.H{"status": "failed", "message": "Falha no pagamento registrada"})
		return

	default:
		log.Printf("â„¹ï¸  Evento nÃ£o tratado: %s", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ignored", "message": "Evento recebido"})
}
