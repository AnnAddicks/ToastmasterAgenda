package main

import (
	"strconv"
	"time"
)

const formatSlashes = "/"
const formatPeriods = "."

func nextTuesday(t time.Time) time.Time {
	const tuesday = 2
	t = t.AddDate(0, 0, (tuesday+(7-int(t.Weekday())))%7)

	return t
}

func formatDate(t time.Time, format string) string {
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	year := strconv.Itoa(t.Year())

	return month + format + day + format + year
}

func agendaMonthDayYear(t time.Time) string {
	day := strconv.Itoa(t.Day())
	month := t.Month().String()
	year := strconv.Itoa(t.Year())

	return month + " " + day + ", " + year
}

func addMinutes(curTime time.Time, minToAdd int) (time.Time, string) {
	nextTime := time.Minute * time.Duration(minToAdd)
	curTime = curTime.Add(nextTime)
	hour, min, _ := curTime.Clock()

	return curTime, strconv.Itoa(hour) + ":" + strconv.Itoa(min)
}
