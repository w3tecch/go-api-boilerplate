package testhelpers

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// RunMiddleware runs a single middleware and returns
// the request and response elements to verify the results
func RunMiddleware(middleware gin.HandlerFunc) (httptest.ResponseRecorder, http.Request) {
	router := gin.New()
	router.Use(middleware)
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "OK")
	})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, r)
	return *w, *r
}
