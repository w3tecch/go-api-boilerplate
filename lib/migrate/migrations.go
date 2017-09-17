package migrate

import "errors"

// ErrorNoNextMigration ...
var ErrorNoNextMigration = errors.New("there is no next version")

// Migrations slice.
type Migrations []*Migration

// Next gets the next migration.
func (ms Migrations) Next(current int64) (*Migration, error) {
	for i, migration := range ms {
		if migration.Version > current {
			return ms[i], nil
		}
	}
	return nil, ErrorNoNextMigration
}
