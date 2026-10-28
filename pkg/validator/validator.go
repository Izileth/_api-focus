package validator

import (
	"net/mail"
	"strings"
)

// IsValidEmail verifies if the provided string is a valid email address
func IsValidEmail(email string) bool {
	if strings.TrimSpace(email) == "" {
		return false
	}
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return addr.Address == email && strings.Contains(email, "@") && strings.Contains(email, ".")
}

// IsValidCurrency checks if the currency is a supported 3-letter ISO code
func IsValidCurrency(currency string) bool {
	c := strings.ToUpper(strings.TrimSpace(currency))
	supported := map[string]bool{
		"BRL": true,
		"USD": true,
		"EUR": true,
	}
	return supported[c]
}

// IsValidAmount verifies that an amount (in cents) is positive
func IsValidAmount(amount int64) bool {
	return amount > 0
}
