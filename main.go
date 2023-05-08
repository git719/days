// main.go

// TODO: Check back on 03:14:08 UTC on 19 January 2038, to confirm we're good ;-)
//       https://en.wikipedia.org/wiki/Year_2038_problem

package main

import (
	"fmt"
	"github.com/git719/utl"
	"os"
)

const (
	// Global constants
	prgname = "days"
	prgver  = "1.0.1"
)

func printUsage() {
	fmt.Printf(prgname + " Calendar days calculator v" + prgver + "\n" +
		"     -DAYS                    Print the date from number of DAYS ago, e.g. -11\n" +
		"     +DAYS                    Print what the date will be in number of DAYS into the future, e.g. +6\n" +
		"     YYYY-MM-DD               Print number of days since given date, or into the future\n" +
		"     YYYY-MM-DD YYYY-MM-DD    Print number of days between the 2 dates\n")
	os.Exit(0)
}

func main() {
	numberOfArguments := len(os.Args[1:]) // Not including the program itself
	if numberOfArguments < 1 || numberOfArguments > 2 {
		// Don't accept less than 1 or more than 2 arguments
		printUsage()
	}

	// Process given arguments
	switch numberOfArguments {
	case 1:
		arg1 := os.Args[1]
		if arg1[0:1] == "+" || arg1[0:1] == "-" {
			dateStr := utl.GetDateInDays(arg1)
			//fmt.Println(dateStr) // Standard format: 2006-01-02 13:50:14 -0500 EST
			fmt.Println(dateStr.Format("2006-01-02"))
		} else if utl.ValidDate(arg1, "2006-01-02") {
			days := utl.GetDaysSinceOrTo(arg1)
			utl.PrintDays(days)
		}
	case 2:
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		if utl.ValidDate(arg1, "2006-01-02") && utl.ValidDate(arg2, "2006-01-02") {
			days := utl.GetDaysBetween(arg1, arg2)
			utl.PrintDays(days)
		}
	default:
		printUsage()
	}
}
