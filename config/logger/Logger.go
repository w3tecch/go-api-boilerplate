package logger

import "github.com/sirupsen/logrus"
import "fmt"
import "github.com/w3tecch/go-api-boilerplate/config/env"

// Logger ...
type Logger struct {
	Scope  string
	Fields logrus.Fields
}

// SetScope ...
func (l *Logger) SetScope(scope string) {
	l.Scope = scope
}

// LineBreak ...
func (l *Logger) LineBreak() {
	fmt.Println()
}

// WithFields ...
func (l *Logger) WithFields(fields logrus.Fields) *Logger {
	newLogger := Logger{Scope: l.Scope, Fields: fields}
	return &newLogger
}

// Debug ...
func (l *Logger) Debug(text string, args ...interface{}) {
	if env.Get().Name == "test" {
		return
	}
	if len(args) == 0 {
		logrus.WithFields(l.Fields).Debug(l.scopePrefix() + text)
	} else {
		logrus.WithFields(l.Fields).Debug(l.scopePrefix()+text, args)
	}
}

// Info ...
func (l *Logger) Info(text string, args ...interface{}) {
	if env.Get().Name == "test" {
		return
	}
	if len(args) == 0 {
		logrus.WithFields(l.Fields).Info(l.scopePrefix() + text)
	} else {
		logrus.WithFields(l.Fields).Info(l.scopePrefix()+text, args)
	}
}

// Warn ...
func (l *Logger) Warn(text string, args ...interface{}) {
	if env.Get().Name == "test" {
		return
	}
	if len(args) == 0 {
		logrus.WithFields(l.Fields).Warn(l.scopePrefix() + text)
	} else {
		logrus.WithFields(l.Fields).Warn(l.scopePrefix()+text, args)
	}
}

// Error ...
func (l *Logger) Error(text string, args ...interface{}) {
	if env.Get().Name == "test" {
		return
	}
	if len(args) == 0 {
		logrus.WithFields(l.Fields).Error(l.scopePrefix() + text)
	} else {
		logrus.WithFields(l.Fields).Error(l.scopePrefix()+text, args)
	}
}

// Fatal ...
func (l *Logger) Fatal(text string, args ...interface{}) {
	if env.Get().Name == "test" {
		return
	}
	if len(args) == 0 {
		logrus.WithFields(l.Fields).Fatal(l.scopePrefix() + text)
	} else {
		logrus.WithFields(l.Fields).Fatal(l.scopePrefix()+text, args)
	}
}

// Panic ...
func (l *Logger) Panic(text string, args ...interface{}) {
	if env.Get().Name == "test" {
		return
	}
	if len(args) == 0 {
		logrus.WithFields(l.Fields).Panic(l.scopePrefix() + text)
	} else {
		logrus.WithFields(l.Fields).Panic(l.scopePrefix()+text, args)
	}
}

////////////////////////////////////////////////////////////
func (l *Logger) scopePrefix() string {
	return "(" + l.Scope + ") "
}
