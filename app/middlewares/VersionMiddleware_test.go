package middlewares

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load("../.env.testing")
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestVersionMiddleware(t *testing.T) {
	os.Setenv("API_VERSION", "0.0.0")
	router := gin.New()
	router.Use(Version()) //only use the mw I want to test
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, r)

	version := w.Header().Get("X-API-VERSION")

	if version == "" {
		t.Log("Failed, because the context or exception is not defined")
		t.Fail()
	}

}
