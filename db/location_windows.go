package db

import (
	"os"
	"path/filepath"
)

func getLocation() string {
	appData := os.Getenv("APPDATA")
	return filepath.Join(appData, "timy", "timy.db")
}
