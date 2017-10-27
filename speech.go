package main

import (
	"errors"
	"strconv"
	"strings"
)

// Speech represents the speech that will be performed for the agenda.
type speech struct {
	manualCode  string
	manualName  string
	details     speechDetails
	displayName string
}

// Factory function to create a speech using a manual and the speech number in that manual.
func (s speech) new(manualCode string, number int) (speech, error) {
	sp := speech{manualCode: manualCode, number: number}

	manual := manualMap[manualCode]

	if number < 1 || number > len(manual.speeches) {
		return speech{}, errors.New("The speech number is not valid for this manual.")
	}

	sp.manualName = manual.manualName
	sp.name = manual.speeches[number-1].name
	sp.min = manual.speeches[number-1].min
	sp.max = manual.speeches[number-1].min
	sp.displayName = "#" + strconv.Itoa(speech.number) + " " + speech.name +
		" " + "(" + strconv.Itoa(speech.min) + "-" + strconv.Itoa(speech.max) + " mins)"
	return sp, nil
}

// Each Toastmaster manual contains 5-10 ordered and named speeches.
// The fields min and max are the minimum and maximum speaking times for a speech.
type speechDetails struct {
	number int
	name   string
	min    int
	max    int
}

// There are 16 Toastmaster manuals a speaker can use at a meeting.
type manual struct {
	manualName string
	speeches   []speechDetails
}

var manualMap = map[string]manual{
	"cc": manual{
		manualName: "Competent Communicator",
		speeches: []speechDetails{
			speechDetails{number: 1, name: "Ice Breaker", min: 4, max: 6},
			speechDetails{number: 2, name: "Organize Your Speech", min: 5, max: 7},
			speechDetails{number: 3, name: "Get to the Point", min: 5, max: 7},
			speechDetails{number: 4, name: "How to Say It", min: 5, max: 7},
			speechDetails{number: 5, name: "Your Body Speaks", min: 5, max: 7},
			speechDetails{number: 6, name: "Vocal Variety", min: 5, max: 7},
			speechDetails{number: 7, name: "Research Your Topic", min: 5, max: 7},
			speechDetails{number: 8, name: "Visual Aids", min: 5, max: 7},
			speechDetails{number: 9, name: "Persuade with Power", min: 5, max: 7},
			speechDetails{number: 10, name: "Inspire Your Audience", min: 8, max: 10},
		},
	},
	"inform": manual{
		manualName: "Speaking to Inform",
		speeches: []speechDetails{
			speechDetails{number: 1, name: "The Speech to Inform", min: 5, max: 7},
			speechDetails{number: 2, name: "Resources for Informing", min: 5, max: 7},
			speechDetails{number: 3, name: "The Demonstration Talk", min: 5, max: 7},
			speechDetails{number: 4, name: "A Fact Finding Report", min: 5, max: 7},
			speechDetails{number: 5, name: "The Abstract Concept", min: 6, max: 8},
		},
	},
	"interpretive": manual{
		manualName: "Interpretive Reading",
		speeches: []speechDetails{
			speechDetails{number: 1, name: "Read a Story", min: 8, max: 10},
			speechDetails{number: 2, name: "Interpretive Poetry", min: 6, max: 8},
			speechDetails{number: 3, name: "The Monodrama", min: 5, max: 7},
			speechDetails{number: 4, name: "The Play", min: 12, max: 15},
			speechDetails{number: 5, name: "The Oratorical Speech", min: 8, max: 10},
		},
	},
	"mgt": manual{
		manualName: "Speeches By Management",
		speeches: []speechDetails{
			speechDetails{number: 1, name: "The Briefing", min: 8, max: 10},
			speechDetails{number: 2, name: "The Technical Speech", min: 8, max: 10},
			speechDetails{number: 3, name: "Manage and Motivate", min: 10, max: 12},
			speechDetails{number: 4, name: "The Status Report", min: 10, max: 12},
			speechDetails{number: 5, name: "Confrontation", min: 5, max: 15},
		},
	},
	"technical": manual{
		manualName: "Technical Presentations",
		speeches: []speechDetails{
			speechDetails{number: 1, name: "Technical Briefing", min: 8, max: 10},
			speechDetails{number: 2, name: "The Proposal", min: 8, max: 10},
			speechDetails{number: 3, name: "The Nontechnical Audience", min: 10, max: 12},
			speechDetails{number: 4, name: "Presenting a Technical Paper", min: 10, max: 12},
			speechDetails{number: 5, name: "The Team Technical Presentation", min: 20, max: 30},
		},
	},
	"storytelling": manual{
		manualName: "Storytelling",
		speeches: []speechDetails{
			speechDetails{number: 1, name: "The Folk Tale", min: 7, max: 9},
			speechDetails{number: 2, name: "Let's Get Personal", min: 6, max: 8},
			speechDetails{number: 3, name: "The Moral of the Story", min: 4, max: 6},
			speechDetails{number: 4, name: "The Touching Story", min: 6, max: 8},
			speechDetails{number: 5, name: "Bringing History to Life", min: 7, max: 9},
		},
	},
}
