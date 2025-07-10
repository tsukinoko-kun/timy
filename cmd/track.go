package cmd

import (
	"errors"
	"time"

	"github.com/spf13/cobra"
	"github.com/tsukinoko-kun/timy/db"
)

var trackCmd = &cobra.Command{
	Use:     "track",
	Short:   "Track time at a day",
	Example: `timy track --date 2023-01-01 --time 1h30m`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var date time.Time
		if cmd.Flag("date").Changed {
			dateStr := cmd.Flag("date").Value.String()
			var err error
			date, err = time.Parse("2006-01-02", dateStr)
			if err != nil {
				return err
			}
		} else {
			date = time.Now()
		}

		var simeSpan time.Duration
		if cmd.Flag("time").Changed {
			simeSpanStr := cmd.Flag("time").Value.String()
			var err error
			simeSpan, err = time.ParseDuration(simeSpanStr)
			if err != nil {
				return err
			}
		} else {
			return errors.New("time is required")
		}

		if err := db.Q.AddTime(
			cmd.Context(),
			simeSpan.String(),
			cmd.Flag("description").Value.String(),
			int64(date.Local().Year()),
			int64(date.Local().Month()),
			int64(date.Local().Day()),
		); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	trackCmd.Flags().StringP("date", "d", "", "Date to track time")
	trackCmd.Flags().StringP("time", "t", "", "Time to track")
	trackCmd.Flags().StringP("description", "m", "", "Description of time")
	trackCmd.MarkFlagRequired("time")
	rootCmd.AddCommand(trackCmd)
}
