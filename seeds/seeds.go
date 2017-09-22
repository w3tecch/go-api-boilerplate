package seeds

import "github.com/w3tecch/go-api-boilerplate/config/logger"

var SeedsLog = logger.Logger{Scope: "seeds"}

func DatabaseSeeds() {
	SeedsLog.Info("Starting seeding")
	SeedUsers()
	SeedsLog.Info("Finished seeding")
}
