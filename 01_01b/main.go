package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	// panic("NOT IMPLEMENTED")
	parsedTimed, err := time.Parse(expectedFormat, target)
	if err != nil { log.Fatal("ERROR: " + err.Error()) }

	// If the target time is before the current time, throw an error
	if time.Now().After(parsedTimed) {
		log.Fatal("ERROR: provided time is before the current time")
	}

	return parsedTimed
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	
	// Get the duration from current time upto to the specified target
	duration := time.Until(target)

	// Convert the nanoseconds into days.
	daysUntilNext := duration.Hours() / 24
	return daysUntilNext
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
