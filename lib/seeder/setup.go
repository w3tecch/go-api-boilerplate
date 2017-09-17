package seeder

import "github.com/w3tecch/go-api-boilerplate/config"

// IsSeederTableReady ...
func IsSeederTableReady() bool {
	db := config.GetDatabaseConnection()
	return db.HasTable("seeder")
}

// CreateSeederTable ...
func CreateSeederTable() (err error) {
	db := config.GetDatabaseConnection()
	err = db.Exec(createSeederTableSQL()).Error
	return err
}

// LockDatabase ...
func LockDatabase() (err error) {
	db := config.GetDatabaseConnection()
	err = db.Exec(getLockStatment(), 1).Error
	return err
}

func createSeederTableSQL() string {
	return `CREATE TABLE seeder (
		is_locked tinyint(1) NOT NULL DEFAULT 1,
		locked_at timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
}

func getLockStatment() string {
	return `INSERT INTO seeder (is_locked) VALUES (?);`
}
