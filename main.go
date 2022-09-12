// main.go

// TODO: Check back on 19th Jan 2038 to see if this passes Millenium bug :-)

package main

import (
	"fmt"
	"os"
)

const (
	// Global constants
	prgname = "days"
	prgver  = "1"
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
			printDateInDays(arg1)
		} else if isDate(arg1) {
		    printDaysSinceOrTo(arg1)
		}
	case 2:
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		if isDate(arg1) && isDate(arg2) {
			printDaysBetween(arg1, arg2)
		}
    default:
		printUsage()
	}
}
