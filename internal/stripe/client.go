package stripeclient

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v78"
)

func Init() {
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		log.Fatal("âš ï¸  STRIPE_SECRET_KEY nÃ£o configurada no .env")
	}

	stripe.Key = key
	log.Println("âœ… Cliente Stripe inicializado!")
}