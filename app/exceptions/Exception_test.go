package exceptions

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestFactoryFunction(t *testing.T) {
	a := NewException(404, "Not Found", errors.New("Test Error"))

	if a.Code != 404 {
		t.Log("Failed, because code is not 404")
		t.Fail()
	}
	if a.Message != "Not Found" {
		t.Log("Failed, because message is not 'Not Found'")
		t.Fail()
	}
}

func TestNewContextPrinter(t *testing.T) {
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	e := NewException(404, "Not Found", errors.New("Test Error"))
	p := e.NewContextPrinter(context)

	if p.Exception.Code != 404 {
		t.Log("Failed, because the context or exception is not defined")
		t.Fail()
	}
}

func TestExceptionPrinter(t *testing.T) {
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	e := NewException(404, "Not Found", errors.New("Test Error"))
	p := e.NewContextPrinter(context)
	p.Print()

	if !p.Context.IsAborted() {
		t.Log("Failed, because it did no abort")
		t.Fail()
	}
}
