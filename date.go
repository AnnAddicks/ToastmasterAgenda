package main

import (
	"strconv"
	"time"
)

const delimiterSlashes = "/"
const delimiterPeriods = "."

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

	return month + delimiter + day + delimiter + year
}

// Formatdate takes in a time and returns a full string month day year with a comma after day.  Ex:  January 1, 2010
func monthDayCommaYear(t time.Time) string {
	day := strconv.Itoa(t.Day())
	month := t.Month().String()
	year := strconv.Itoa(t.Year())

	return month + " " + day + ", " + year
}

// addMinutes takes in a time and minutes to add to that time and returns the new time and a string representation (ex:  10:44).
func addMinutes(curTime time.Time, minToAdd int) (time.Time, string) {
	nextTime := time.Minute * time.Duration(minToAdd)
	curTime = curTime.Add(nextTime)
	hour, min, _ := curTime.Clock()

	return curTime, strconv.Itoa(hour) + ":" + strconv.Itoa(min)
}
