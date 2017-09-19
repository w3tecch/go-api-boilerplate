package seeds

import (
	"github.com/w3tecch/go-api-boilerplate/config"
	"github.com/w3tecch/go-api-boilerplate/lib/logger"
)

// SeedsLog ...
var SeedsLog = logger.Logger{Scope: "seeds"}

// SeedUsers ...
func SeedUsers() {
	SeedsLog.Info("Starting seeding users")

	db := config.GetDatabaseConnection()
	for i := 1; i <= 100; i++ {
		user := FakeUser()
		err := db.Create(&user).Error
		if err != nil {
			SeedsLog.Error("Failed seeding user", err.Error())
			return
		}
	}

	SeedsLog.Info("Finished seeding users")
}
