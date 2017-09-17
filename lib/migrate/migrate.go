package migrate

import (
	"github.com/sirupsen/logrus"

	"github.com/w3tecch/go-api-boilerplate/lib/logger"
)

// MigrationLog ...
var MigrationLog = logger.Logger{Scope: "migrate"}

// Sync ...
func Sync() {
	MigrationLog.Info("Start migrating database")

	// Check if migration table is setup
	if !IsMigrationTableReady() {
		MigrationLog.Info("Creating migration table in the database")
		if err := CreateMigrationTable(); err != nil {
			MigrationLog.Fatal("Failed to create migration table")
		} else {
			MigrationLog.Info("Successfully created migration table")
		}
	}

	// Get Version of the database
	version, err := GetVersionOfDatabase()
	if err != nil {
		MigrationLog.Fatal("Failed to read the latest migrtion of the database")
	}
	MigrationLog.WithFields(logrus.Fields{"version": version}).Info("Read current migration version of the database")

	// Collect all migrations
	ms, err := CollectMigrationScripts("./migrations", version)
	if err != nil {
		MigrationLog.Fatal("Failed to collect the migration files")
	} else {
		MigrationLog.WithFields(logrus.Fields{"amount": len(ms)}).Info("Found open migrations")
		for _, mig := range ms {
			MigrationLog.Info("--> " + mig.Name)
		}
	}

	// Run all new migrations
	err = Run(ms)
	if err != nil {
		MigrationLog.WithFields(logrus.Fields{"error": err.Error()}).Fatal("Could not migrate the database")
	} else {
		MigrationLog.Info("Database was successfully migrated")
	}

}
