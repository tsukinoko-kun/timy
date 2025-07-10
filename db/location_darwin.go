package db

import (
	"os"
	"path/filepath"
)

func getLocation() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, "Library", "Application Support", "timy", "timy.db")
}
