package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/config/env"
)

// Version is a middleware function that appends the Drone version information
// to the HTTP response. This is intended for debugging and troubleshooting.
func Version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-API-VERSION", env.Get().APIVersion)
		c.Next()
	}
}
