package main

import (
	"testing"
	"time"
)

//Create a static date that is a Monday
func testDate() time.Time {
	d := "2017-01-02 7:11"
	t, _ := time.Parse("2006-01-02 15:04", d)

	return t
}

func TestgetNextTuesday(t *testing.T) {
	d := testDate()
	c := getNextTuesday(d)

	tuesday := d.AddDate(0, 0, 1)
	if c != tuesday {
		t.Error("Expected '1.3.2017', got ", c)
	}
}

func TestAgendaDate(t *testing.T) {
	d := testDate()
	c := AgendaDate(d)

	if c != "1.2.2017" {
		t.Error("Expected '1.2.2017', got ", c)
	}
}

func TestDateWithSlashes(t *testing.T) {
	d := testDate()
	c := DateWithSlashes(d)

	if c != "1/2/2017" {
		t.Error("Expected '1/2/2017', got ", c)
	}
}

func TestAgendaMonthDayYear(t *testing.T) {
	d := testDate()
	c := AgendaMonthDayYear(d)

	if c != "January 2, 2017" {
		t.Error("Expected 'January 2, 2017', got ", c)
	}

}

func TestPrettyPrintTime(t *testing.T) {
	n := testDate()
	m := 1

	_, test := prettyPrintTime(n, m)
	if test != "7:12" {
		t.Error("Expected an empty string, got: ", test)
	}

}

func TestPrettyPrintTimeIncrementHour(t *testing.T) {
	n := testDate()
	m := 60

	_, test := prettyPrintTime(n, m)
	if test != "8:11" {
		t.Error("Expected an empty string, got: ", test)
	}

}

func TestPrettyPrintTimeNextTime(t *testing.T) {
	n := testDate()
	m := 60

	newTime, _ := prettyPrintTime(n, m)
	if newTime.Equal(n) {
		t.Error("Expected the dates to be equal, but got: ", newTime)
	}
}
