package seeder

import (
	"fmt"

	"github.com/w3tecch/go-api-boilerplate/config/database"
)

// IsSeederTableReady ...
func IsSeederTableReady() bool {
	db := database.Connection()
	return db.HasTable("seeder")
}

// CreateSeederTable ...
func CreateSeederTable() (err error) {
	db := database.Connection()
	err = db.Exec(createSeederTableSQL()).Error
	return err
}

// LockDatabase ...
func LockDatabase() (err error) {
	db := database.Connection()
	err = db.Exec(updateLockStatment()).Error
	fmt.Println(updateLockStatment())
	fmt.Println(err)
	return err
}

// InsertLockDatabase ...
func InsertLockDatabase() (err error) {
	db := database.Connection()
	err = db.Exec(getInsertLockStatment(), 1).Error
	return err
}

func createSeederTableSQL() string {
	return `CREATE TABLE seeder (
		is_locked tinyint(1) NOT NULL DEFAULT 1,
		locked_at timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
}

func getInsertLockStatment() string {
	return `INSERT INTO seeder (is_locked) VALUES (?);`
}

func updateLockStatment() string {
	return `UPDATE seeder SET is_locked=1, locked_at=CURRENT_TIMESTAMP WHERE is_locked=0;`
}
