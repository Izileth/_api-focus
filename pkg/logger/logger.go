package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// Level defines log severity
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger provides structured logging functions
type Logger struct {
	mu     sync.Mutex
	level  Level
	logger *log.Logger
}

var defaultLogger = New(os.Stdout, LevelInfo)

// New creates a new Logger instance
func New(out io.Writer, level Level) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(out, "", log.LstdFlags),
	}
}

// SetLevel updates the logger minimum severity level
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// Log writes formatted log if the level is above threshold
func (l *Logger) Log(level Level, msg string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < l.level {
		return
	}
	prefix := fmt.Sprintf("[%s] ", level.String())
	formatted := fmt.Sprintf(msg, args...)
	_ = l.logger.Output(2, prefix+formatted)
}

// Debug logs debug messages
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Log(LevelDebug, msg, args...)
}

// Info logs informative messages
func (l *Logger) Info(msg string, args ...interface{}) {
	l.Log(LevelInfo, msg, args...)
}

// Warn logs warning messages
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Log(LevelWarn, msg, args...)
}

// Error logs error messages
func (l *Logger) Error(msg string, args ...interface{}) {
	l.Log(LevelError, msg, args...)
}

// Global convenience wrappers
func Debug(msg string, args ...interface{}) {
	defaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	defaultLogger.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	defaultLogger.Error(msg, args...)
}
