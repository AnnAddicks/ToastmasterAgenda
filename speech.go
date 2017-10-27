package main

import (
	"strconv"
	"strings"
)

type Speech struct {
	manualCode string
	manualName string
	name       string
	number     int
	min        int
	max        int
}

func GetSpeech(manualName string, number int) Speech {
	speech := Speech{number: number, manualCode: manualName}

	switch strings.ToLower(manualName) {
	case "cc":
		speech.manualName = "Competent Communicator"

		switch number {
		case 1:
			speech.name = "Ice Breaker"
			speech.min = 4
			speech.max = 6

		case 2:
			speech.name = "Organize Your Speech"
			speech.min = 5
			speech.max = 7

		case 3:
			speech.name = "Get to the Point"
			speech.min = 5
			speech.max = 7

		case 4:
			speech.name = "How to Say It"
			speech.min = 5
			speech.max = 7

		case 5:
			speech.name = "Your Body Speaks"
			speech.min = 5
			speech.max = 7

		case 6:
			speech.name = "Vocal Variety"
			speech.min = 5
			speech.max = 7

		case 7:
			speech.name = "Research Your Topic"
			speech.min = 5
			speech.max = 7

		case 8:
			speech.name = "Visual Aids"
			speech.min = 5
			speech.max = 7

		case 9:
			speech.name = "Persuade with Power"
			speech.min = 5
			speech.max = 7

		case 10:
			speech.name = "Inspire Your Audience"
			speech.min = 8
			speech.max = 10

		}

	case "inform":
		speech.manualName = "Speaking to Inform"
		switch number {
		case 1:
			speech.name = "The Speech to Inform"
			speech.min = 5
			speech.max = 7

		case 2:
			speech.name = "Resources for Informing"
			speech.min = 5
			speech.max = 7

		case 3:
			speech.name = "The Demonstration Tak"
			speech.min = 5
			speech.max = 7

		case 4:
			speech.name = "A Fact Finding Report"
			speech.min = 5
			speech.max = 7

		case 5:
			speech.name = "The Abstract Concept"
			speech.min = 6
			speech.max = 8
		}

	case "interpretive":
		speech.manualName = "Interpretive Reading"
		switch number {
		case 1:
			speech.name = "Read a Story"
			speech.min = 8
			speech.max = 10

		case 2:
			speech.name = "Interpretive Poetry"
			speech.min = 6
			speech.max = 8

		case 3:
			speech.name = "The Monodrama"
			speech.min = 5
			speech.max = 7

		case 4:
			speech.name = "The Play"
			speech.min = 12
			speech.max = 15

		case 5:
			speech.name = "The Oratorical Speech"
			speech.min = 8
			speech.max = 10
		}
	case "mgt":
		speech.manualName = "Speeches by Management "
		switch number {
		case 1:
			speech.name = "The Briefing"
			speech.min = 8
			speech.max = 10

		case 2:
			speech.name = "The Technical Speech"
			speech.min = 8
			speech.max = 10

		case 3:
			speech.name = "Manage and Motivate"
			speech.min = 10
			speech.max = 12

		case 4:
			speech.name = "The Status Report"
			speech.min = 10
			speech.max = 12

		case 5:
			speech.name = "Confrontation"
			speech.min = 5
			speech.max = 15
		}

	case "technical":
		speech.manualName = "Technical Presentations"
		switch number {
		case 1:
			speech.name = "Technical Briefing"
			speech.min = 8
			speech.max = 10

		case 2:
			speech.name = "The Proposal (Add Q&A TIME!!!!)"
			speech.min = 8
			speech.max = 10

		case 3:
			speech.name = "The Nontechnical Audience"
			speech.min = 10
			speech.max = 12

		case 4:
			speech.name = "Presenting a Technical Paper"
			speech.min = 10
			speech.max = 12

		case 5:
			speech.name = "The Team Technical Presentation"
			speech.min = 20
			speech.max = 30
		}
	case "storytelling":
		speech.manualName = "Storytelling"
		switch number {
		case 1:
			speech.name = "The Folk Tale"
			speech.min = 7
			speech.max = 9

		case 2:
			speech.name = "Let's Get Personal"
			speech.min = 6
			speech.max = 8

		case 3:
			speech.name = "The Moral of the Story"
			speech.min = 4
			speech.max = 6

		case 4:
			speech.name = "The Touching Story"
			speech.min = 6
			speech.max = 8

		case 5:
			speech.name = "Bringing History to Life"
			speech.min = 7
			speech.max = 9
		}
	}

	speech.name = "#" + strconv.Itoa(speech.number) + " " + speech.name +
		" " + "(" + strconv.Itoa(speech.min) + "-" + strconv.Itoa(speech.max) + " mins)"
	return speech
}
