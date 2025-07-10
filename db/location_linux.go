package db

import (
	"os"
	"path/filepath"
)

func getLocation() string {
	if xdg, ok := os.LookupEnv("XDG_DATA_HOME"); ok {
		return filepath.Join(xdg, "timy", "timy.db")
	}
	home := os.Getenv("HOME")
	return filepath.Join(home, ".local", "share", "timy", "timy.db")
}
