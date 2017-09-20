package seeder

import (
	"time"

	"github.com/w3tecch/go-api-boilerplate/config"
)

// SeederLog ...
var SeederLog = config.Logger{Scope: "seeder"}

// SeederRecord ...
type SeederRecord struct {
	IsLocked   *bool
	MigratedAt *time.Time
}

// TableName set MigrationRecord's table name to be `version`
func (SeederRecord) TableName() string {
	return "seeder"
}

// Sync ...
func Sync() {

	// Check if migration table is setup
	if !IsSeederTableReady() {
		SeederLog.Info("Creating seeding table in the database")
		if err := CreateSeederTable(); err != nil {
			SeederLog.Fatal("Failed to create seeding table")
		} else {
			SeederLog.Info("Successfully created seeding table")
			if err := InsertLockDatabase(); err != nil {
				SeederLog.Fatal("Failed to lock seeding table")
			} else {
				SeederLog.Info("Successfully locked seeding table")
			}
		}
	}

	// Give a short feedback if seeding is allowed. Only the
	// database can tell us that
	if IsSeedingAllowed() {
		SeederLog.Info("Seeding is allowed")
	} else {
		SeederLog.Info("Seeding is not allowed")
	}

}

// IsSeedingAllowed ...
func IsSeedingAllowed() bool {
	db := config.GetDatabaseConnection()
	sr := SeederRecord{}

	err := db.Last(&sr).Error
	if err != nil {
		return false
	}

	return !*sr.IsLocked
}
