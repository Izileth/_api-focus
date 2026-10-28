package error

import (
	"errors"
	"net/http"
	"testing"
)

func TestAppError(t *testing.T) {
	appErr := New(http.StatusBadRequest, "invalid request")
	if appErr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, appErr.Code)
	}
	if appErr.Error() != "invalid request" {
		t.Errorf("expected message 'invalid request', got '%s'", appErr.Error())
	}

	rootErr := errors.New("database connection failed")
	wrapped := Wrap(http.StatusInternalServerError, "internal server error", rootErr)
	if wrapped.Unwrap() != rootErr {
		t.Errorf("expected unwrapped error to match rootErr")
	}
	expectedMsg := "internal server error: database connection failed"
	if wrapped.Error() != expectedMsg {
		t.Errorf("expected '%s', got '%s'", expectedMsg, wrapped.Error())
	}
}

func TestErrorHelpers(t *testing.T) {
	br := BadRequest("bad request")
	if br.Code != http.StatusBadRequest || br.Message != "bad request" {
		t.Errorf("unexpected BadRequest error: %+v", br)
	}

	nf := NotFound("not found")
	if nf.Code != http.StatusNotFound || nf.Message != "not found" {
		t.Errorf("unexpected NotFound error: %+v", nf)
	}

	ua := Unauthorized("unauthorized")
	if ua.Code != http.StatusUnauthorized || ua.Message != "unauthorized" {
		t.Errorf("unexpected Unauthorized error: %+v", ua)
	}

	ie := Internal("internal error", errors.New("crash"))
	if ie.Code != http.StatusInternalServerError || ie.Err == nil {
		t.Errorf("unexpected Internal error: %+v", ie)
	}
}
