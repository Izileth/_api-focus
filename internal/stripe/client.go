package stripeclient

import (
	"os"

	"github.com/stripe/stripe-go/v78"
)

func Init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}