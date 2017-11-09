package main

import (
	"testing"
)

func TestSpeechNew(t *testing.T) {
	s := NewSpeech("CC", 10)

	if s.number != 10 {
		t.Error("Expected '10', got ", s.number)
	}

	if s.manualCode != "cc" {
		t.Error("Expected 'cc', got", s.manualCode)
	}

	if s.manualName != "Competent Communicator" {
		t.Error("Expected 'Competent Communicator', got", s.manualName)
	}

	if s.name != "Inspire Your Audience" {
		t.Error("Expected 'Inspire Your Audience', got", s.name)
	}

	if s.min != 8 {
		t.Error("Expected '8', got", s.min)
	}

	if s.max != 10 {
		t.Error("Expected '10', got", s.max)
	}
}

func TestSpeechNewWithInvalidNumber(t *testing.T) {
	s := NewSpeech("CC", 11)

	if s.number != 11 {
		t.Error("Expected '11', got ", s.number)
	}

	if s.manualCode != "cc" {
		t.Error("Expected 'cc', got", s.manualCode)
	}

	if s.manualName != "Competent Communicator" {
		t.Error("Expected 'Competent Communicator', got", s.manualName)
	}

	if s.name != "" {
		t.Error("Expected '', got", s.name)
	}

	if s.min != 0 {
		t.Error("Expected '0', got", s.min)
	}

	if s.max != 0 {
		t.Error("Expected '0', got", s.max)
	}
}

func TestSpeechInfo(t *testing.T) {
	s := NewSpeech("CC", 1)

	if s.info() != "#1 Ice Breaker (4-6 mins)" {
		t.Error("Expected '#1 Ice Breaker (4-6 mins)', got", s.info())
	}

}
