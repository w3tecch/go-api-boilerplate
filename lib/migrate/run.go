package migrate

// Run ...
func Run(ms Migrations) error {
	for {
		// Check if the database is already up to date
		current := currentVersion()
		next, err := ms.Next(current)
		if err != nil {
			if err == ErrorNoNextMigration {
				MigrationLog.Info("Database is up to data")
				return nil
			}
			return err
		}

		// Runs the up migrations and if it fails it will automatically
		// run the down migration to have a clean version in the database
		if errUp := next.Up(); errUp != nil {
			if errDown := next.Down(); errDown != nil {
				return errDown
			}
			return errUp
		}

	}
}

func currentVersion() int64 {
	version, err := GetVersionOfDatabase()
	if err != nil {
		MigrationLog.Fatal("Failed to read the latest migrtion of the database")
		return 9223372036854775807 // retrun max of int64
	}
	return version
}
