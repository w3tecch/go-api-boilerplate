package testhelpers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/app/routes"
	"github.com/w3tecch/go-api-boilerplate/config/database"
	"github.com/w3tecch/go-api-boilerplate/lib/migrate"
)

// SetupServer ...
func SetupServer() *gin.Engine {
	os.Setenv("ENVIRONMENT", "test")
	os.Setenv("DB_DIALECT", "mysql")
	os.Setenv("API_TITLE", "title")
	os.Setenv("API_VERSION", "0.0.0")
	os.Setenv("DB_CONNECTION", "root@tcp(localhost:3306)/boilerplate-test?charset=utf8&parseTime=True")
	os.Setenv("AUTH0_BASE_URL", "http://localhost:3333")

	gin.SetMode(gin.TestMode)
	migrate.Sync()
	router := routes.Router()
	return router
}

func TearDownServer() {
	database.Clear()
	os.Exit(0)
}
