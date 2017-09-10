package ctrl

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// APIInfo ...
type APIInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// GetAPIInfo ...
func GetAPIInfo(c *gin.Context) {
	// res := lib.Response{ResponseWriter: w}
	// res.SendOK(APIInfo{
	// 	Name:    os.Getenv("API_TITLE"),
	// 	Version: os.Getenv("API_VERSION"),
	// })
	c.JSON(http.StatusOK, APIInfo{
		Name:    os.Getenv("API_TITLE"),
		Version: os.Getenv("API_VERSION"),
	})
}
