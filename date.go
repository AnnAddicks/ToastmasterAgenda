package main

import (
	"strconv"
	"time"
)

const delimiterSlashes = "/"
const delimiterPeriods = "."

var availableDelimiters = map[string]bool{delimiterSlashes: true, delimiterPeriods: true}

// NextTuesday takes in a date and returns that date if it is Tuesday or the following Tuesday.
func nextTuesday(t time.Time) time.Time {
	const tuesday = 2
	t = t.AddDate(0, 0, (tuesday+(7-int(t.Weekday())))%7)

	return t
}

// Formatdate takes in a time and returns a numeric month day year with a delimiter in between each.
func formatDate(t time.Time, delimiter string) string {
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	year := strconv.Itoa(t.Year())

	d := delimiterSlashes
	if availableDelimiters[delimiter] {
		d = delimiter
	}
	return month + d + day + d + year
}

// MonthDayCommaYear takes in a time and returns a full string month day year with a comma after day
// (ex:  January 1, 2010).
func monthDayCommaYear(t time.Time) string {
	day := strconv.Itoa(t.Day())
	month := t.Month().String()
	year := strconv.Itoa(t.Year())

	return month + " " + day + ", " + year
}

// AddMinutes takes in a time and minutes to add to that time and returns the new time and a string
// representation (ex:  10:44).
func addMinutes(curTime time.Time, minToAdd int) (time.Time, string) {
	nextTime := time.Minute * time.Duration(minToAdd)
	curTime = curTime.Add(nextTime)
	hour, min, _ := curTime.Clock()

	return curTime, strconv.Itoa(hour) + ":" + strconv.Itoa(min)
}
