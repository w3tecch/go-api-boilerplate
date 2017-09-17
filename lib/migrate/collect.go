package migrate

import (
	"fmt"
	"os"
	"path/filepath"
)

var registeredGoMigrations = map[int64]*Migration{}

// CollectMigrationScripts ...
func CollectMigrationScripts(dirpath string, currentVersion int64) (ms Migrations, err error) {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s directory does not exists", dirpath)
	}

	// SQL migration files.
	sqlMigrationFiles, err := filepath.Glob(dirpath + "/**.sql")
	if err != nil {
		return nil, err
	}
	for _, file := range sqlMigrationFiles {
		v, err := ParseTimestamp(file)
		if err != nil {
			return nil, err
		}
		f, err := ParseName(file)
		if err != nil {
			return nil, err
		}
		if versionFilter(v, currentVersion) {
			migration := &Migration{Name: f, Version: v, Source: file}
			ms = append(ms, migration)
		}
	}

	return ms, nil
}

func versionFilter(v int64, current int64) bool {
	return v > current
}
