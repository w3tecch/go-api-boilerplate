package middlewares

import (
	"fmt"
	"net/http"

	"github.com/w3tecch/go-api-boilerplate/app/lib"
)

// SaySomething ...
func SaySomething() lib.Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("before wuhuuu middleware")
			h.ServeHTTP(w, r)
			fmt.Println("after wuhuuu middleware")
		}
	}
}
