package main

import (
	"testing"
)


func TestParseManualAndNumber(t *testing.T) {

	speech := "Ann Addicks\nCC #4 "

	name, manual, num := parseManualAndNumber(speech)

	if name != "Ann Addicks" {
		t.Error("Expected 'Ann Addicks', got ", name)
	}

	if manual != "CC" {
		t.Error("Expected 'CC', got ", manual)
	}

	if num != 4 {
		t.Error("Expected '4', got ", num)
	}
}