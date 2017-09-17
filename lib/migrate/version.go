package migrate

import (
	"github.com/w3tecch/go-api-boilerplate/config"
)

// GetVersionOfDatabase returns the latest migration
func GetVersionOfDatabase() (version int64, err error) {
	db := config.GetDatabaseConnection()
	mr := MigrationRecord{}

	err = db.Last(&mr).Error
	// Only error is that the record was not found
	if err != nil {
		return 0, nil
	}

	version, err = ParseTimestamp(*mr.Name)
	if err != nil {
		return 0, err
	}

	return version, err
}

// InsertVersionSQL ...
func InsertVersionSQL() string {
	return "INSERT INTO version (name) VALUES (?);"
}

// DeleteVersionSQL ...
func DeleteVersionSQL() string {
	return "DELETE FROM version WHERE name = ?;"
}
