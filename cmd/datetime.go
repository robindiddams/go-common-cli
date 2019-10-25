package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pinpt/go-common/datetime"
	"github.com/spf13/cobra"
)

var datetimeCmd = &cobra.Command{
	Use:   "datetime",
	Short: "datetime package",
}

func parseStringEpoch(str string) (int64, bool) {
	ts, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, false
	}
	return ts, true
}

func stringEpochToDate(str string) (time.Time, bool) {
	ts, ok := parseStringEpoch(str)
	if !ok {
		return time.Time{}, false
	}
	return datetime.DateFromEpoch(ts), true
}

var datetimeDateRangeCmd = &cobra.Command{
	Use:   "DateRange",
	Short: "datetime.DateRange(time.Time, int64)",
	Long: `Will run datetime.DateRange.
Examples:

datetime DateRange 1571282055407 alltime
datetime DateRange 1571282055407 30`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		reftimeRaw := args[0]
		var timeunit int64
		if args[1] == "alltime" {
			timeunit = -1
		} else {
			var err error
			timeunit, err = strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Println("bad time unit")
				os.Exit(1)
			}
		}
		var reftime time.Time
		reftime, ok := stringEpochToDate(reftimeRaw)
		if !ok {
			// TODO: parse isodate
		}
		start, end := datetime.DateRange(reftime, timeunit)
		fmt.Printf("start: %d end: %d\n", start, end)
	},
}

var datetimeEndofDayCmd = &cobra.Command{
	Use:   "EndofDay",
	Short: "datetime.EndofDay(int64)",
	Long: `Will run datetime.EndofDay.
Examples:

gc datetime EndofDay 1571282055407`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reftime, ok := parseStringEpoch(args[0])
		if !ok {
			fmt.Println("unparsable timestamp:", args[0])
			os.Exit(1)
		}
		eod := datetime.EndofDay(reftime)
		fmt.Printf("%d\n", eod)
	},
}

var datetimeStartofDayCmd = &cobra.Command{
	Use:   "StartofDay",
	Short: "datetime.StartofDay(int64)",
	Long: `Will run datetime.StartofDay.
Examples:

gc datetime StartofDay 1571282055407`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reftime, ok := parseStringEpoch(args[0])
		if !ok {
			fmt.Println("unparsable timestamp:", args[0])
			os.Exit(1)
		}
		sod := datetime.StartofDay(reftime)
		fmt.Printf("%d\n", sod)
	},
}

var datetimeEpochNowCmd = &cobra.Command{
	Use:   "EpochNow",
	Short: "datetime.EpochNow()",
	Long: `Will run datetime.EpochNow.
Examples:

gc datetime EpochNow`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%d\n", datetime.EpochNow())
	},
}

func init() {
	rootCmd.AddCommand(datetimeCmd)
	datetimeCmd.AddCommand(datetimeDateRangeCmd)
	datetimeCmd.AddCommand(datetimeEndofDayCmd)
	datetimeCmd.AddCommand(datetimeStartofDayCmd)
	datetimeCmd.AddCommand(datetimeEpochNowCmd)
}
