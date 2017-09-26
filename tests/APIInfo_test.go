package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/w3tecch/go-api-boilerplate/app/controllers"
	"github.com/w3tecch/go-api-boilerplate/lib/testhelpers"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func TestMain(m *testing.M) {
	server = testhelpers.SetupServer()
	m.Run()
	testhelpers.TearDownServer()
}

func TestAPIInfoEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/info", nil)
	server.ServeHTTP(w, r)

	res := controllers.APIInfo{}
	testhelpers.JSON(*w, &res)

	if res.Name != "title" && res.Version != "0.0.0" {
		t.Log("Failed, because env vars are not set")
		t.Fail()
	}

}
