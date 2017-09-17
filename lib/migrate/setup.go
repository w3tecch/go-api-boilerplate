package migrate

import "github.com/w3tecch/go-api-boilerplate/config"

// IsMigrationTableReady ...
func IsMigrationTableReady() bool {
	db := config.GetDatabaseConnection()
	return db.HasTable("version")
}

// CreateMigrationTable ...
func CreateMigrationTable() (err error) {
	db := config.GetDatabaseConnection()
	err = db.Exec(createVersionTableSQL()).Error
	return err
}

func createVersionTableSQL() string {
	return `CREATE TABLE version (
            id int(10) unsigned NOT NULL AUTO_INCREMENT,
            name varchar(255) UNIQUE,
            migrated_at timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
            PRIMARY KEY(id)
        	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
}
