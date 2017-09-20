package seeds

import "github.com/w3tecch/go-api-boilerplate/config"

var SeedsLog = config.Logger{Scope: "seeds"}

func DatabaseSeeds() {
	SeedsLog.Info("Starting seeding")
	SeedUsers()
	SeedsLog.Info("Finished seeding")
}
