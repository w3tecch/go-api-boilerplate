package middlewares

import (
	"os"

	auth0middleware "github.com/w3tecch/go-auth0-middleware"
)

// Auth0Middleware ...
func Auth0Middleware() *auth0middleware.Auth0Middleware {
	return auth0middleware.New(auth0middleware.Options{
		Endpoint:   os.Getenv("AUTH0_BASE_URL"),
		ContextKey: "TokenInfo",
	})
}
