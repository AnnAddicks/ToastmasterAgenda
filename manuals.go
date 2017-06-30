package main

import ()

type Speech struct {
	manualName string
	speechName string
	number     int
	min        int
	max        int
}

func manuals(manualName string, number int) {
	speech := Speech{}
	speech.number = number

	switch manualName {

	case "CC":
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
	}

	return speech
}
