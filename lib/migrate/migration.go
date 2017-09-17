package migrate

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/w3tecch/go-api-boilerplate/config"
)

// MigrationRecord ...
type MigrationRecord struct {
	ID         *int64
	Name       *string
	MigratedAt *time.Time
}

// TableName set MigrationRecord's table name to be `version`
func (MigrationRecord) TableName() string {
	return "version"
}

// Migration struct.
type Migration struct {
	Name    string
	Version int64
	Source  string
}

func (m *Migration) String() string {
	return fmt.Sprintf(m.Name)
}

// Up runs an up migration.
func (m *Migration) Up() (err error) {
	MigrationLog.WithFields(logrus.Fields{
		"name": m.Name,
	}).Info("Run up migration")
	return m.Run(true)
}

// Down runs an down migration.
func (m *Migration) Down() (err error) {
	MigrationLog.WithFields(logrus.Fields{
		"name": m.Name,
	}).Info("Run down migration")
	return m.Run(false)
}

// Run runs a migration.
func (m *Migration) Run(direction bool) (err error) {
	f, err := os.Open(m.Source)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	statements := getSQLStatements(f, direction)
	db := config.GetDatabaseConnection()

	// Begin a transaction
	tx := db.Begin()
	for _, query := range statements {
		if err = tx.Exec(query).Error; err != nil {
			// Rollback the transaction in case of error
			tx.Rollback()
			return err
		}
	}

	// Add migration to version table
	if direction {
		if err = tx.Exec(InsertVersionSQL(), m.Name).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err = tx.Exec(DeleteVersionSQL(), m.Name).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}
