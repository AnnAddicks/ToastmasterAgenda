package main

import (
	"testing"
	"time"
)

//Create a static date that is a Monday.
func testDate() time.Time {
	d := "2017-01-02 7:11"
	t, _ := time.Parse("2006-01-02 15:04", d)

	return t
}

func TestGetNextTuesday(t *testing.T) {
	d := testDate()
	c := nextTuesday(d)

	tuesday := d.AddDate(0, 0, 1)
	if c != tuesday {
		t.Error("Expected '1.3.2017', got ", c)
	}
}

func TestFormatDate(t *testing.T) {
	d := testDate()
	c := formatDate(d, delimiterPeriods)

	if c != "1.2.2017" {
		t.Error("Expected '1.2.2017', got ", c)
	}
}

func TestDateWithSlashes(t *testing.T) {
	d := testDate()
	c := formatDate(d, delimiterSlashes)

	if c != "1/2/2017" {
		t.Error("Expected '1/2/2017', got ", c)
	}
}

func TestMonthDayCommaYear(t *testing.T) {
	d := testDate()
	c := monthDayCommaYear(d)

	if c != "January 2, 2017" {
		t.Error("Expected 'January 2, 2017', got ", c)
	}

}

func TestPrettyPrintTime(t *testing.T) {
	n := testDate()
	m := 1

	_, test := addMinutes(n, m)
	if test != "7:12" {
		t.Error("Expected an empty string, got: ", test)
	}

}

func TestPrettyPrintTimeIncrementHour(t *testing.T) {
	n := testDate()
	m := 60

	_, test := addMinutes(n, m)
	if test != "8:11" {
		t.Error("Expected an empty string, got: ", test)
	}

}

func TestPrettyPrintTimeNextTime(t *testing.T) {
	n := testDate()
	m := 60

	newTime, _ := addMinutes(n, m)
	if newTime.Equal(n) {
		t.Error("Expected the dates to be equal, but got: ", newTime)
	}
}
