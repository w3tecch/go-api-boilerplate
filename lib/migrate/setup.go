package migrate

import "github.com/w3tecch/go-api-boilerplate/config/database"
import "fmt"

// IsMigrationTableReady ...
func IsMigrationTableReady() bool {
	db := database.Connection()
	return db.HasTable("version")
}

// CreateMigrationTable ...
func CreateMigrationTable() (err error) {
	db := database.Connection()
	err = db.Exec(createVersionTableSQL()).Error
	fmt.Println(err)
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
