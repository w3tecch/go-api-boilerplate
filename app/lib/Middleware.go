package lib

import (
	"net/http"
)

// MiddlewareFunc ...
type MiddlewareFunc func(h http.HandlerFunc) http.HandlerFunc

// Middleware ...
type Middleware func(h http.HandlerFunc) http.HandlerFunc

// Use ...
func Use(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
