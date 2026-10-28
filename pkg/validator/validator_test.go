package validator

import "testing"

func TestIsValidEmail(t *testing.T) {
	valid := []string{
		"test@example.com",
		"user.name@domain.co",
		"customer+filter@sub.domain.org",
	}
	for _, email := range valid {
		if !IsValidEmail(email) {
			t.Errorf("expected '%s' to be valid", email)
		}
	}

	invalid := []string{
		"",
		"   ",
		"plainaddress",
		"@missingusername.com",
		"username@.com",
		"user@domain",
	}
	for _, email := range invalid {
		if IsValidEmail(email) {
			t.Errorf("expected '%s' to be invalid", email)
		}
	}
}

func TestIsValidCurrency(t *testing.T) {
	valid := []string{"BRL", "usd", "Eur", "  brl  "}
	for _, c := range valid {
		if !IsValidCurrency(c) {
			t.Errorf("expected '%s' to be valid currency", c)
		}
	}

	invalid := []string{"", "GBP", "JPY", "INVALID", "123"}
	for _, c := range invalid {
		if IsValidCurrency(c) {
			t.Errorf("expected '%s' to be invalid currency", c)
		}
	}
}

func TestIsValidAmount(t *testing.T) {
	if !IsValidAmount(100) {
		t.Errorf("expected 100 to be valid")
	}
	if !IsValidAmount(1) {
		t.Errorf("expected 1 to be valid")
	}
	if IsValidAmount(0) {
		t.Errorf("expected 0 to be invalid")
	}
	if IsValidAmount(-50) {
		t.Errorf("expected -50 to be invalid")
	}
}
