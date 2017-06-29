package main

import (
	"testing"
	"time"
)

func testDate() time.Time {
	d := "2017-01-02"
	t, _ := time.Parse("2006-01-02", d)

	return t
}

func TestAgendaDate(t *testing.T) {
	d := testDate()
	c := AgendaDate(d)

	if c != "1.2.2017" { 
		t.Error("Expected '1.2.2017', got ", c)
	}
}

func TestAgendaMonthDayYear(t *testing.T) {
	d := testDate()
	c := AgendaMonthDayYear(d)

	if c != "January 2, 2017" { 
		t.Error("Expected 'January 2, 2017', got ", c)
	}


}
