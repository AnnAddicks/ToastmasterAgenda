package main

import (
	"testing"
)

func TestSpeechNew(t *testing.T) {
	s := NewSpeech("CC", 10)

	if s.Number != 10 {
		t.Error("Expected '10', got ", s.Number)
	}

	if s.ManualCode != "cc" {
		t.Error("Expected 'cc', got", s.ManualCode)
	}

	if s.ManualName != "Competent Communicator" {
		t.Error("Expected 'Competent Communicator', got", s.ManualName)
	}

	if s.Name != "Inspire Your Audience" {
		t.Error("Expected 'Inspire Your Audience', got", s.Name)
	}

	if s.Min != 8 {
		t.Error("Expected '8', got", s.Min)
	}

	if s.Max != 10 {
		t.Error("Expected '10', got", s.Max)
	}
}

func TestSpeechNewWithInvalidNumber(t *testing.T) {
	s := NewSpeech("CC", 11)

	if s.Number != 11 {
		t.Error("Expected '11', got ", s.Number)
	}

	if s.ManualCode != "cc" {
		t.Error("Expected 'cc', got", s.ManualCode)
	}

	if s.ManualName != "Competent Communicator" {
		t.Error("Expected 'Competent Communicator', got", s.ManualName)
	}

	if s.Name != "" {
		t.Error("Expected '', got", s.Name)
	}

	if s.Min != 0 {
		t.Error("Expected '0', got", s.Min)
	}

	if s.Max != 0 {
		t.Error("Expected '0', got", s.Max)
	}
}

func TestSpeechInfo(t *testing.T) {
	s := NewSpeech("CC", 1)

	if s.info() != "#1 Ice Breaker (4-6 mins)" {
		t.Error("Expected '#1 Ice Breaker (4-6 mins)', got", s.info())
	}

}
