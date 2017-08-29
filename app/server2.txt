package app

import (
	"fmt"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

// NewServer ...
func NewServer() *negroni.Negroni {

	// Define the global middlewares
	server := negroni.New()
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(middlewares.CORSMiddleware())
	server.Use(middlewares.SecureMiddleware())
	server.Use(middlewares.LogMiddleware())

	// auth0Middleware := middlewares.Auth0Middleware{Endpoint: ""}
	auth0Middleware := new(middlewares.Auth0Middleware)
	fmt.Println(auth0Middleware)
	server.Use(auth0Middleware)

	// Attach app router
	server.UseHandler(NewRouter())

	return server
}
