package seeds

import "github.com/w3tecch/go-api-boilerplate/lib/logger"

// SeedsLog ...
var SeedsLog = logger.Logger{Scope: "seeds"}

// SeedUsers ...
func SeedUsers() {
	SeedsLog.LineBreak()
	SeedsLog.Info("Starting seeding users")

}
