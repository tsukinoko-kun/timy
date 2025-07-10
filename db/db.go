package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var Q *Queries

func init() {
	filename := getLocation()
	_ = os.MkdirAll(filepath.Dir(filename), 0755)
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec(`PRAGMA journal_mode = WAL`); err != nil {
		panic(err)
	}

	if err := Migrate(context.Background(), db); err != nil {
		panic(err)
	}

	Q = New(db)
}

func Close() {
	if Q == nil {
		return
	}
	if db, ok := Q.db.(*sql.DB); ok {
		if err := db.Close(); err != nil {
			panic(err)
		}
		Q = nil
	}
}
