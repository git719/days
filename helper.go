// helper.go

package main

import (
	"fmt"
	"time"
	"strconv"
)

func intAbs(val int64) int64 {
	if val < 0 {
		return -val
	}
	return val
}

func stringToInt(s string) int64 {
	// Convert string number to INT64
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}	

func stringToTime(s string) (time.Time, error) {
	// Convert Unix epoc seconds string to Time type
    sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
    }
	return time.Unix(sec, 0), nil
}

func stringDateToEpoc(s string) (int64, error) {
	// Convert data string to Unix epoc seconds
	t, err := time.Parse("2006-01-02", s) // First, convert string to Time
	if err != nil {
		panic(err.Error())
	}
	return t.Unix(), nil  // Finally, convert Time type to Unix epoc seconds
}

func isDate(s string) bool {
	// Check if string is a valid yyyy-mm-dd date format
	_, err := time.Parse("2006-01-02", s)
	if err != nil {
		return false
	}
	return true
}

func printDateInDays(days string) {
	// Print yyyy-mm-dd date for given number of +/- days in future or past
	now := time.Now().Unix()
	now = now + (stringToInt(days) * 86400) // 86400 seconds in a day
	nowStr := strconv.FormatInt(now, 10)
	date1, err := stringToTime(nowStr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(date1.Format("2006-01-02"))
}

func printDays(days int64) {
	days_abs := intAbs(days)
	if days_abs > 365 {
		years := days_abs / 365
		modulus := days_abs % 365
		fmt.Printf("%d (%d years + %d days)\n", days, years, modulus)
	} else {
		fmt.Println(days)
	}
}

func printDaysSinceOrTo(date1 string) {
	// Calculate and print number of +/- days from NOW to date given
	// Note: Calculations are all in UTC time
	epoc1, err := stringDateToEpoc(date1)
	if err != nil {
		panic(err.Error())
	}

	now := time.Now()
	now_secs := now.Unix()

	// Get today's epoc time at midnight UTC
	today := now.Format("2006-01-02")
	t, err := time.Parse("2006-01-02", today) // First, convert string to Time
	if err != nil {
		panic(err.Error())
	}
	midnight_secs := t.Unix()

	since_midnight := now_secs - midnight_secs
    epoc1 = epoc1 + since_midnight
	printDays( (epoc1 - now_secs) / 86400 )
}

func printDaysBetween(date1, date2 string) {
	// Calculate and print number of days between 2 dates
	epoc1, err := stringDateToEpoc(date1)
	if err != nil {
		panic(err.Error())
	}
	epoc2, err := stringDateToEpoc(date2)
	if err != nil {
		panic(err.Error())
	}
	if epoc1 > epoc2 {
		printDays( (epoc1 - epoc2) / 86400 )
	} else if epoc2 > epoc1 {
		printDays( (epoc2 - epoc1) / 86400 )
	} else {
		fmt.Println(0)		
	}
}
