package migrate

import (
	"errors"
	"path/filepath"
	"strconv"
	"strings"
)

// ParseTimestamp validate the filepath and returns the timestamp of it
func ParseTimestamp(name string) (int64, error) {

	base := filepath.Base(name)

	if ext := filepath.Ext(base); ext != ".sql" {
		return 0, errors.New("not a recognized migration file type")
	}

	idx := strings.Index(base, "_")
	if idx < 0 {
		return 0, errors.New("no separator found")
	}

	n, e := strconv.ParseInt(base[:idx], 10, 64)
	if e == nil && n <= 0 {
		return 0, errors.New("migration IDs must be greater than zero")
	}

	return n, e
}

// ParseName validate the filepath and returns the name of it
func ParseName(name string) (string, error) {
	base := filepath.Base(name)
	return base, nil
}
