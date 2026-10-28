package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestVersionMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name            string
		version         string
		expectedVersion string
	}{
		{
			name:            "v1 version header",
			version:         "v1",
			expectedVersion: "v1",
		},
		{
			name:            "v2 version header",
			version:         "v2",
			expectedVersion: "v2",
		},
		{
			name:            "v3 version header",
			version:         "v3",
			expectedVersion: "v3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(VersionMiddleware(tt.version))
			r.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("expected status 200, got %d", w.Code)
			}

			if got := w.Header().Get("X-API-Version"); got != tt.expectedVersion {
				t.Errorf("expected X-API-Version %s, got %s", tt.expectedVersion, got)
			}
		})
	}
}
