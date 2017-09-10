package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Version is a middleware function that appends the Drone version information
// to the HTTP response. This is intended for debugging and troubleshooting.
func Version(c *gin.Context) {
	c.Header("X-API-VERSION", os.Getenv("API_VERSION"))
	c.Next()
}
