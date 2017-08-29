package app

import (
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
	"github.com/w3tecch/go-api-boilerplate/app/middlewares"
)

// NewServer ...
func NewServer() *negroni.Negroni {

	// Define the global middlewares
	server := negroni.New()
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(middlewares.CORSMiddleware())
	server.Use(middlewares.SecureMiddleware())
	server.Use(middlewares.Auth0Middleware())
	server.Use(middlewares.LogMiddleware())

	// Attach app router
	server.UseHandler(NewRouter())

	return server
}
