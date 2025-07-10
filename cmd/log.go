package cmd

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/tsukinoko-kun/timy/db"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log time",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			year  *int64
			month *int64
			day   *int64
		)

		if cmd.Flag("year").Changed {
			yearStr := cmd.Flag("year").Value.String()
			yearUsr, err := strconv.ParseInt(yearStr, 10, 64)
			if err != nil {
				return err
			}
			year = &yearUsr
		}

		if cmd.Flag("month").Changed {
			monthStr := cmd.Flag("month").Value.String()
			var err error
			monthUsr, err := strconv.ParseInt(monthStr, 10, 64)
			if err != nil {
				return err
			}
			month = &monthUsr
			if year == nil {
				// default to current year
				year = ptr(int64(time.Now().Local().Year()))
			}
		}

		if cmd.Flag("day").Changed {
			dayStr := cmd.Flag("day").Value.String()
			dayUsr, err := strconv.ParseInt(dayStr, 10, 64)
			if err != nil {
				return err
			}
			day = &dayUsr
			if month == nil {
				// default to current month
				month = ptr(int64(time.Now().Local().Month()))
				if year != nil {
					return errors.New("year is set but month is not, makes no sense")
				} else {
					year = ptr(int64(time.Now().Local().Year()))
				}
			}
		}

		var times []db.GetTimes

		if year != nil {
			if month != nil {
				if day != nil {
					if timesConcrete, err := db.Q.GetTimesYearMonthDay(cmd.Context(), *year, *month, *day); err != nil {
						return err
					} else {
						times = db.ConvertGetTimes(timesConcrete)
					}
				} else {
					if timesConcrete, err := db.Q.GetTimesYearMonth(cmd.Context(), *year, *month); err != nil {
						return err
					} else {
						times = db.ConvertGetTimes(timesConcrete)
					}
				}
			} else {
				if timesConcrete, err := db.Q.GetTimesYear(cmd.Context(), *year); err != nil {
					return err
				} else {
					times = db.ConvertGetTimes(timesConcrete)
				}
			}
		} else {
			return errors.New("no year, month or day is set")
		}

		totalTimespan := time.Duration(0)
		comments := strings.Builder{}
		for _, timespan := range times {
			ts, err := timespan.GetTimespan()
			if err != nil {
				return err
			}
			totalTimespan += ts
			comments.WriteString(timespan.GetDescription())
			comments.WriteString("\n\n")
		}

		cmd.Printf("Total timespan: %s\n", totalTimespan)

		if commentsStr := strings.TrimSpace(comments.String()); commentsStr != "" {
			cmd.Printf("Comments:\n%s", commentsStr)
		}

		return nil
	},
}

func init() {
	logCmd.Flags().StringP("year", "y", "", "Year to log")
	logCmd.Flags().StringP("month", "m", "", "Month to log")
	logCmd.Flags().StringP("day", "d", "", "Day to log")
	rootCmd.AddCommand(logCmd)
}

func ptr[T any](t T) *T {
	return &t
}
