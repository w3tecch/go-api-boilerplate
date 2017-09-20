package controllers

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/lib/seeder"
	"github.com/w3tecch/go-api-boilerplate/seeds"
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

// Seeding ...
func (ctrl APIController) Seeding(c *gin.Context) {
	if !seeder.IsSeedingAllowed() {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	// Seeding data to the database
	seeds.DatabaseSeeds()

	// After seeding the database will be locked again
	seeder.LockDatabase()
	c.JSON(http.StatusOK, gin.H{})
}
