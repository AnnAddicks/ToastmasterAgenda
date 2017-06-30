package main

import (
	"testing"
	"time"
)

//Create a static date that is a Monday
func testDate() time.Time {
	d := "2017-01-02"
	t, _ := time.Parse("2006-01-02", d)

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
