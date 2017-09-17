package controllers

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

// APIInfo ...
type APIInfo struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	GOVersion string `json:"goVersion"`
}

// APIController ...
type APIController struct{}

// GetInfo ...
func (ctrl APIController) GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, APIInfo{
		Name:      os.Getenv("API_TITLE"),
		Version:   os.Getenv("API_VERSION"),
		GOVersion: runtime.Version(),
	})
}
