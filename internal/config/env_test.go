package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	val := GetEnv("TEST_KEY")
	if val != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", val)
	}
}
