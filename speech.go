package main

import (
	"log"
	"strconv"
	"strings"
)

// SpeechDetails contains the information of a speech.
// The fields Min and Max are the minimum and maximum speaking times for a Speech.
type SpeechDetails struct {
	Number int
	Name   string
	Min    int
	Max    int
}

// Speech represents the Speech that will be performed for the agenda.
type Speech struct {
	ManualCode string
	ManualName string
	SpeechDetails
}

// Info creates a string that represents a Speech ex: "#1 Ice Breaker (4-6 mins)."
func (s *Speech) info() string {
	return "#" + strconv.Itoa(s.Number) + " " + s.Name +
		" " + "(" + strconv.Itoa(s.Min) + "-" + strconv.Itoa(s.Max) + " mins)"
}

// NewSpeech is a factory function to create a Speech using a manual code and the Speech number in that manual.
func NewSpeech(manCode string, num int) *Speech {
	manCode = strings.ToLower(manCode)
	man := manualMap[manCode]

	sp := &Speech{
		ManualCode: manCode,
		ManualName: man.manualName,
		SpeechDetails: SpeechDetails{
			Number: num,
		},
	}

	// Do not fail with invalid input, return with the default values set.
	if num < 1 || num > len(man.speeches) {
		log.Print("Speech num is invalid for the man.  Manual code: "+manCode+" Speech num:", num)
		return sp
	}

	sp.Name = man.speeches[num-1].Name
	sp.Min = man.speeches[num-1].Min
	sp.Max = man.speeches[num-1].Max

	return sp
}

// There are 16 Toastmaster manuals a Speaker can use at a meeting.
type manual struct {
	manualName string
	speeches   []SpeechDetails
}

// Representation of the Toasmaters speeches with their code as the key.
var manualMap = map[string]manual{
	"cc": {
		manualName: "Competent Communicator",
		speeches: []SpeechDetails{
			{Number: 1, Name: "Ice Breaker", Min: 4, Max: 6},
			{Number: 2, Name: "Organize Your Speech", Min: 5, Max: 7},
			{Number: 3, Name: "Get to the Point", Min: 5, Max: 7},
			{Number: 4, Name: "How to Say It", Min: 5, Max: 7},
			{Number: 5, Name: "Your Body Speaks", Min: 5, Max: 7},
			{Number: 6, Name: "Vocal Variety", Min: 5, Max: 7},
			{Number: 7, Name: "Research Your Topic", Min: 5, Max: 7},
			{Number: 8, Name: "Visual Aids", Min: 5, Max: 7},
			{Number: 9, Name: "Persuade with Power", Min: 5, Max: 7},
			{Number: 10, Name: "Inspire Your Audience", Min: 8, Max: 10},
		},
	},
	"inform": {
		manualName: "Speaking to Inform",
		speeches: []SpeechDetails{
			{Number: 1, Name: "The Speech to Inform", Min: 5, Max: 7},
			{Number: 2, Name: "Resources for Informing", Min: 5, Max: 7},
			{Number: 3, Name: "The Demonstration Talk", Min: 5, Max: 7},
			{Number: 4, Name: "A Fact Finding Report", Min: 5, Max: 7},
			{Number: 5, Name: "The Abstract Concept", Min: 6, Max: 8},
		},
	},
	"interpretive": {
		manualName: "Interpretive Reading",
		speeches: []SpeechDetails{
			{Number: 1, Name: "Read a Story", Min: 8, Max: 10},
			{Number: 2, Name: "Interpretive Poetry", Min: 6, Max: 8},
			{Number: 3, Name: "The Monodrama", Min: 5, Max: 7},
			{Number: 4, Name: "The Play", Min: 12, Max: 15},
			{Number: 5, Name: "The Oratorical Speech", Min: 8, Max: 10},
		},
	},
	"mgt": {
		manualName: "Speeches By Management",
		speeches: []SpeechDetails{
			{Number: 1, Name: "The Briefing", Min: 8, Max: 10},
			{Number: 2, Name: "The Technical Speech", Min: 8, Max: 10},
			{Number: 3, Name: "Manage and Motivate", Min: 10, Max: 12},
			{Number: 4, Name: "The Status Report", Min: 10, Max: 12},
			{Number: 5, Name: "Confrontation", Min: 5, Max: 15},
		},
	},
	"technical": {
		manualName: "Technical Presentations",
		speeches: []SpeechDetails{
			{Number: 1, Name: "Technical Briefing", Min: 8, Max: 10},
			{Number: 2, Name: "The Proposal", Min: 8, Max: 10},
			{Number: 3, Name: "The Nontechnical Audience", Min: 10, Max: 12},
			{Number: 4, Name: "Presenting a Technical Paper", Min: 10, Max: 12},
			{Number: 5, Name: "The Team Technical Presentation", Min: 20, Max: 30},
		},
	},
	"storytelling": {
		manualName: "Storytelling",
		speeches: []SpeechDetails{
			{Number: 1, Name: "The Folk Tale", Min: 7, Max: 9},
			{Number: 2, Name: "Let's Get Personal", Min: 6, Max: 8},
			{Number: 3, Name: "The Moral of the Story", Min: 4, Max: 6},
			{Number: 4, Name: "The Touching Story", Min: 6, Max: 8},
			{Number: 5, Name: "Bringing History to Life", Min: 7, Max: 9},
		},
	},
	"specialty": {
		manualName: "Specialty Speeches",
		speeches: []SpeechDetails{
			{Number: 1, Name: "Impromptu Speaking", Min: 5, Max: 7},
			{Number: 2, Name: "Uplift the Spirit", Min: 8, Max: 10},
			{Number: 3, Name: "Sell a product", Min: 10, Max: 12},
			{Number: 4, Name: "Read Out Loud", Min: 12, Max: 15},
			{Number: 5, Name: "Introduce the Speaker", Min: 0, Max: 0},
		},
	},
	"hpl": {
		manualName: "High Performance Leadership",
		speeches: []SpeechDetails{
			{Number: 1, Name: "Share Your Vision", Min: 5, Max: 6},
			{Number: 2, Name: "Presenting The Results", Min: 5, Max: 7},
		},
	},
}
