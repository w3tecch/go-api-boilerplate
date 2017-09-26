package middlewares

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/w3tecch/go-api-boilerplate/lib/testhelpers"
)

func TestMain(m *testing.M) {
	godotenv.Load("../.env.testing")
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestVersionMiddleware(t *testing.T) {
	os.Setenv("API_VERSION", "0.0.0")
	w, _ := testhelpers.RunMiddleware(Version())
	version := w.Header().Get("X-API-VERSION")

	if version == "" {
		t.Log("Failed, because the context or exception is not defined")
		t.Fail()
	}
}
