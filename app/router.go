package app

import (
	"github.com/w3tecch/go-api-boilerplate/app/ctrl"
	"github.com/w3tecch/go-api-boilerplate/app/middlewares"

	"github.com/gin-gonic/gin"
)

// Router ...
func Router() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// ------------------------------
	// Logger middleware will write the logs to gin.DefaultWriter even you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Custom middlewares
	r.Use(middlewares.Secure)
	r.Use(middlewares.CORS)
	r.Use(middlewares.NoCache)
	r.Use(middlewares.Version)

	// Controller Routes
	// ------------------------------
	// Return the api information to the user
	r.GET("/api/info", ctrl.GetAPIInfo)

	return r
}
