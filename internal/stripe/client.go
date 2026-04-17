package stripeclient

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkoutsession"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func Init() {
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		log.Fatal("⚠️ STRIPE_SECRET_KEY não configurada no .env")
	}

	stripe.Key = key
	log.Println("✅ Cliente Stripe inicializado!")
}

// CreatePixIntent cria uma intenção de pagamento via Pix
func CreatePixIntent(amount int64, currency, email, name, taxID, line1, city, state, postalCode string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"pix"}),
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Type: stripe.String("pix"),
			BillingDetails: &stripe.PaymentIntentPaymentMethodDataBillingDetailsParams{
				Email: stripe.String(email),
				Name:  stripe.String(name),
				Address: &stripe.AddressParams{
					Line1:      stripe.String(line1),
					City:       stripe.String(city),
					State:      stripe.String(state),
					PostalCode: stripe.String(postalCode),
					Country:    stripe.String("BR"),
				},
			},
		},
	}
	return paymentintent.New(params)
}

// CreateBoletoIntent cria uma intenção de pagamento via Boleto
func CreateBoletoIntent(amount int64, currency, email, name, taxID, line1, city, state, postalCode string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"boleto"}),
		PaymentMethodData: &stripe.PaymentIntentPaymentMethodDataParams{
			Type: stripe.String("boleto"),
			BillingDetails: &stripe.PaymentIntentPaymentMethodDataBillingDetailsParams{
				Email: stripe.String(email),
				Name:  stripe.String(name),
				Address: &stripe.AddressParams{
					Line1:      stripe.String(line1),
					City:       stripe.String(city),
					State:      stripe.String(state),
					PostalCode: stripe.String(postalCode),
					Country:    stripe.String("BR"),
				},
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

// CreateCardIntent cria uma intenção de pagamento via Cartão
func CreateCardIntent(amount int64, currency string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}
	return paymentintent.New(params)
}

// CreateCheckoutSession cria uma sessão de checkout e retorna a URL para pagamento
func CreateCheckoutSession(amount int64, currency, successURL, cancelURL string) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
			"pix",
			"boleto",
		}),
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
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL),
		CancelURL:  stripe.String(cancelURL),
	}
	return checkoutsession.New(params)
}
