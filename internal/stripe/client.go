package stripeclient

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func Init() {
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		log.Fatal("âš ï¸  STRIPE_SECRET_KEY nÃ£o configurada no .env")
	}

	stripe.Key = key
	log.Println("âœ… Cliente Stripe inicializado!")
}

// CreatePixIntent cria uma intenÃ§Ã£o de pagamento via Pix
func CreatePixIntent(amount int64, currency, email, name string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"pix"}),
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Type: stripe.String("pix"),
			BillingDetails: &stripe.PaymentIntentPaymentMethodDataBillingDetailsParams{
				Email: stripe.String(email),
				Name:  stripe.String(name),
			},
		},
	}
	return paymentintent.New(params)
}

// CreateBoletoIntent cria uma intenÃ§Ã£o de pagamento via Boleto
func CreateBoletoIntent(amount int64, currency, email, name, taxID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"boleto"}),
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Type: stripe.String("boleto"),
			BillingDetails: &stripe.PaymentIntentPaymentMethodDataBillingDetailsParams{
				Email: stripe.String(email),
				Name:  stripe.String(name),
			},
			Boleto: &stripe.PaymentMethodBoletoParams{
				TaxID: stripe.String(taxID),
			},
		},
		PaymentMethodOptions: &stripe.PaymentIntentPaymentMethodOptionsParams{
			Boleto: &stripe.PaymentIntentPaymentMethodOptionsBoletoParams{
				ExpiresAfterDays: stripe.Int64(3),
			},
		},
	}
	return paymentintent.New(params)
}

// CreateCardIntent cria uma intenÃ§Ã£o de pagamento via CartÃ£o
func CreateCardIntent(amount int64, currency string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}
	return paymentintent.New(params)
}