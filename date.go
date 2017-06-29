package main

import (
	"strconv"
	"time"
)

func getNextTuesday() time.Time {
	t := time.Now()
	for int(t.Weekday()) != 2 { //todo figure out how to use the constant instead of special number 2! t.Weekday.Tuesday {
		t = t.AddDate(0, 0, 1)
	}

	return t
}

func AgendaDate(t time.Time) string {
	month := t.Month().String()
	day := strconv.Itoa(t.Day())
	year := strconv.Itoa(t.Year())

	return month + "." + day + "." + year
}

func AgendaMonthDayYear(t time.Time) string {
	day := strconv.Itoa(t.Day())
	month := strconv.Itoa(int(t.Month()))
	year := strconv.Itoa(t.Year())

	return month + "/" + day + "/" + year
}
