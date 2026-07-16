package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRootInfo(t *testing.T) {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup router
	r := gin.New()
	r.GET("/", RootInfo)

	// Create request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Validate fields
	if response["message"] != "Welcome to API Focus" {
		t.Errorf("Expected message 'Welcome to API Focus', got '%v'", response["message"])
	}

	versions, ok := response["versions"].([]interface{})
	if !ok {
		t.Fatalf("Expected 'versions' to be a slice, got %T", response["versions"])
	}

	expectedVersions := []string{"v1", "v2", "v3"}
	if len(versions) != len(expectedVersions) {
		t.Errorf("Expected %d versions, got %d", len(expectedVersions), len(versions))
	}

	for i, v := range versions {
		if v != expectedVersions[i] {
			t.Errorf("Expected version %s at index %d, got %s", expectedVersions[i], i, v)
		}
	}
}

func TestHealthCheck_DBNil(t *testing.T) {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup router
	r := gin.New()
	r.GET("/health", HealthCheck)

	// Create request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Parse response
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Validate fields
	if response["api"] != "UP" {
		t.Errorf("Expected api status 'UP', got '%v'", response["api"])
	}

	if response["database"] != "DOWN" {
		t.Errorf("Expected database status 'DOWN' when DB is nil, got '%v'", response["database"])
	}

	if response["message"] != "API Focus systems check" {
		t.Errorf("Expected message 'API Focus systems check', got '%v'", response["message"])
	}
}
