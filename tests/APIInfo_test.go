package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/w3tecch/go-api-boilerplate/app/controllers"
	"github.com/w3tecch/go-api-boilerplate/lib/migrate"
	"github.com/w3tecch/go-api-boilerplate/lib/testhelpers"

	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/app/routes"
	"github.com/w3tecch/go-api-boilerplate/config/database"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	os.Setenv("ENVIRONMENT", "test")
	os.Setenv("DB_DIALECT", "mysql")
	os.Setenv("API_TITLE", "title")
	os.Setenv("API_VERSION", "0.0.0")
	os.Setenv("DB_CONNECTION", "root@tcp(localhost:3306)/boilerplate-test?charset=utf8&parseTime=True")
	os.Setenv("AUTH0_BASE_URL", "http://localhost:3333")

	gin.SetMode(gin.TestMode)
	migrate.Sync()
	router = routes.Router()

	m.Run()

	// Teardown section
	database.Clear()
	os.Exit(0)
}

func TestAPIInfoEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/info", nil)
	router.ServeHTTP(w, r)

	res := controllers.APIInfo{}
	testhelpers.JSON(*w, &res)

	if res.Name != "title" && res.Version != "0.0.0" {
		t.Log("Failed, because env vars are not set")
		t.Fail()
	}

}
