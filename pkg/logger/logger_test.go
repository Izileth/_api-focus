package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	l := New(&buf, LevelInfo)

	l.Debug("debug message %d", 1)
	if strings.Contains(buf.String(), "[DEBUG]") {
		t.Errorf("Debug log should have been filtered out")
	}

	l.Info("hello %s", "world")
	if !strings.Contains(buf.String(), "[INFO] hello world") {
		t.Errorf("Expected info message in buffer, got '%s'", buf.String())
	}

	buf.Reset()
	l.Warn("warning %d", 42)
	if !strings.Contains(buf.String(), "[WARN] warning 42") {
		t.Errorf("Expected warn message in buffer, got '%s'", buf.String())
	}

	buf.Reset()
	l.Error("failure %s", "boom")
	if !strings.Contains(buf.String(), "[ERROR] failure boom") {
		t.Errorf("Expected error message in buffer, got '%s'", buf.String())
	}
}

func TestLoggerSetLevel(t *testing.T) {
	var buf bytes.Buffer
	l := New(&buf, LevelError)

	l.Warn("should be skipped")
	if buf.Len() > 0 {
		t.Errorf("Expected empty buffer, got '%s'", buf.String())
	}

	l.SetLevel(LevelDebug)
	l.Debug("now visible")
	if !strings.Contains(buf.String(), "[DEBUG] now visible") {
		t.Errorf("Expected debug message after SetLevel, got '%s'", buf.String())
	}
}

func TestLevelString(t *testing.T) {
	if LevelDebug.String() != "DEBUG" {
		t.Errorf("expected DEBUG, got %s", LevelDebug.String())
	}
	if LevelInfo.String() != "INFO" {
		t.Errorf("expected INFO, got %s", LevelInfo.String())
	}
	if LevelWarn.String() != "WARN" {
		t.Errorf("expected WARN, got %s", LevelWarn.String())
	}
	if LevelError.String() != "ERROR" {
		t.Errorf("expected ERROR, got %s", LevelError.String())
	}
	if Level(999).String() != "UNKNOWN" {
		t.Errorf("expected UNKNOWN, got %s", Level(999).String())
	}
}
