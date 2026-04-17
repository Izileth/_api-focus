package stripeclient

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func Init() {
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		log.Fatal("⚠️ STRIPE_SECRET_KEY não configurada")
	}

	stripe.Key = key
	log.Println("✅ Stripe inicializado")
}

// ✅ PaymentIntent moderno (Stripe decide método: PIX, cartão, boleto)
func CreatePaymentIntent(amount int64, currency, email string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),

		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},

		ReceiptEmail: stripe.String(email),
	}

	return paymentintent.New(params)
}

// ✅ Checkout Session (sem forçar métodos)
func CreateCheckoutSession(amount int64, currency, successURL, cancelURL string) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),

		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Pagamento Focus"),
					},
					UnitAmount: stripe.Int64(amount),
				},
				Quantity: stripe.Int64(1),
			},
		},

		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
	}

	return session.New(params)
}