package seeds

import (
	"github.com/w3tecch/go-api-boilerplate/config"
)

// SeedUsers ...
func SeedUsers() {
	db := config.GetDatabaseConnection()
	for i := 1; i <= 100; i++ {
		user := FakeUser()
		err := db.Create(&user).Error
		if err != nil {
			SeedsLog.Error("Failed seeding user", err.Error())
			return
		}
	}
}
