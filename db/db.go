package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

var Q *Queries

func init() {
	filename := getLocation()
	_ = os.MkdirAll(filepath.Dir(filename), 0755)
	db, err := sql.Open("sqlite", "file:"+filename)
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

func (ymd GetTimesYearMonthDayRow) GetTimespan() (time.Duration, error) {
	return time.ParseDuration(ymd.Timespan)
}

func (ymd GetTimesYearMonthDayRow) GetDescription() string {
	return ymd.Description
}

func (ym GetTimesYearMonthRow) GetTimespan() (time.Duration, error) {
	return time.ParseDuration(ym.Timespan)
}

func (ym GetTimesYearMonthRow) GetDescription() string {
	return ym.Description
}

func (y GetTimesYearRow) GetTimespan() (time.Duration, error) {
	return time.ParseDuration(y.Timespan)
}

func (y GetTimesYearRow) GetDescription() string {
	return y.Description
}

type GetTimes interface {
	GetTimespan() (time.Duration, error)
	GetDescription() string
}

func ConvertGetTimes[T GetTimes](times []T) []GetTimes {
	converted := make([]GetTimes, len(times))
	for i, t := range times {
		converted[i] = t
	}
	return converted
}
