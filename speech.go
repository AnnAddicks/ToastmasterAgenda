package main

import (
	"log"
	"strconv"
	"strings"
)

// Each Toastmaster manual contains 5-10 ordered and named speeches.
// The fields min and max are the minimum and maximum speaking times for a Speech.
type SpeechDetails struct {
	number int
	name   string
	min    int
	max    int
}

// Speech represents the Speech that will be performed for the agenda.
type Speech struct {
	manualCode string
	manualName string
	SpeechDetails
}

// Info creates a string that represents a Speech ex: "#1 Ice Breaker (4-6 mins)."
func (s Speech) info() string {
	return "#" + strconv.Itoa(s.number) + " " + s.name +
		" " + "(" + strconv.Itoa(s.min) + "-" + strconv.Itoa(s.max) + " mins)"
}

// Factory function to create a Speech using a manual code and the Speech number in that manual.
func NewSpeech(manCode string, num int) Speech {
	manCode = strings.ToLower(manCode)
	man := manualMap[manCode]

	sp := Speech{
		manualCode: manCode,
		manualName: man.manualName,
		SpeechDetails: SpeechDetails{
			number: num,
		},
	}

	// Do not fail with invalid input, return with the default values set.
	if num < 1 || num > len(man.speeches) {
		log.Print("Speech num is invalid for the man.  Manual code: "+manCode+" Speech num:", num)
		return sp
	}

	sp.name = man.speeches[num-1].name
	sp.min = man.speeches[num-1].min
	sp.max = man.speeches[num-1].max

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
			{number: 1, name: "Ice Breaker", min: 4, max: 6},
			{number: 2, name: "Organize Your Speech", min: 5, max: 7},
			{number: 3, name: "Get to the Point", min: 5, max: 7},
			{number: 4, name: "How to Say It", min: 5, max: 7},
			{number: 5, name: "Your Body Speaks", min: 5, max: 7},
			{number: 6, name: "Vocal Variety", min: 5, max: 7},
			{number: 7, name: "Research Your Topic", min: 5, max: 7},
			{number: 8, name: "Visual Aids", min: 5, max: 7},
			{number: 9, name: "Persuade with Power", min: 5, max: 7},
			{number: 10, name: "Inspire Your Audience", min: 8, max: 10},
		},
	},
	"inform": {
		manualName: "Speaking to Inform",
		speeches: []SpeechDetails{
			{number: 1, name: "The Speech to Inform", min: 5, max: 7},
			{number: 2, name: "Resources for Informing", min: 5, max: 7},
			{number: 3, name: "The Demonstration Talk", min: 5, max: 7},
			{number: 4, name: "A Fact Finding Report", min: 5, max: 7},
			{number: 5, name: "The Abstract Concept", min: 6, max: 8},
		},
	},
	"interpretive": {
		manualName: "Interpretive Reading",
		speeches: []SpeechDetails{
			{number: 1, name: "Read a Story", min: 8, max: 10},
			{number: 2, name: "Interpretive Poetry", min: 6, max: 8},
			{number: 3, name: "The Monodrama", min: 5, max: 7},
			{number: 4, name: "The Play", min: 12, max: 15},
			{number: 5, name: "The Oratorical Speech", min: 8, max: 10},
		},
	},
	"mgt": {
		manualName: "Speeches By Management",
		speeches: []SpeechDetails{
			{number: 1, name: "The Briefing", min: 8, max: 10},
			{number: 2, name: "The Technical Speech", min: 8, max: 10},
			{number: 3, name: "Manage and Motivate", min: 10, max: 12},
			{number: 4, name: "The Status Report", min: 10, max: 12},
			{number: 5, name: "Confrontation", min: 5, max: 15},
		},
	},
	"technical": {
		manualName: "Technical Presentations",
		speeches: []SpeechDetails{
			{number: 1, name: "Technical Briefing", min: 8, max: 10},
			{number: 2, name: "The Proposal", min: 8, max: 10},
			{number: 3, name: "The Nontechnical Audience", min: 10, max: 12},
			{number: 4, name: "Presenting a Technical Paper", min: 10, max: 12},
			{number: 5, name: "The Team Technical Presentation", min: 20, max: 30},
		},
	},
	"storytelling": {
		manualName: "Storytelling",
		speeches: []SpeechDetails{
			{number: 1, name: "The Folk Tale", min: 7, max: 9},
			{number: 2, name: "Let's Get Personal", min: 6, max: 8},
			{number: 3, name: "The Moral of the Story", min: 4, max: 6},
			{number: 4, name: "The Touching Story", min: 6, max: 8},
			{number: 5, name: "Bringing History to Life", min: 7, max: 9},
		},
	},
	"specialty": {
		manualName: "Specialty Speeches",
		speeches: []SpeechDetails{
			{number: 1, name: "Impromptu Speaking", min: 5, max: 7},
			{number: 2, name: "Uplift the Spirit", min: 8, max: 10},
			{number: 3, name: "Sell a product", min: 10, max: 12},
			{number: 4, name: "Read Out Loud", min: 12, max: 15},
			{number: 5, name: "Introduce the Speaker", min: 0, max: 0},
		},
	},
	"hpl": {
		manualName: "High Performance Leadership",
		speeches: []SpeechDetails{
			{number: 1, name: "Share Your Vision", min: 5, max: 6},
			{number: 2, name: "Presenting The Results", min: 5, max: 7},
		},
	},
}
