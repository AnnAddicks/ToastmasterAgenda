package main

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

// Represents the board members of a Toastmasters meeting.
type board struct {
	president string
	vpe       string
	vpm       string
	vppr      string
	secretary string
	treasurer string
	saa       string
}

// Factory method using a spreadsheet to fill in board members.
func (board) new(sheet *spreadsheet.Sheet) board {
	board := board{}
	board.president = sheet.Columns[1][0].Value
	board.vpe = sheet.Columns[1][1].Value
	board.vpm = sheet.Columns[1][2].Value
	board.vppr = sheet.Columns[1][3].Value
	board.secretary = sheet.Columns[1][4].Value
	board.treasurer = sheet.Columns[1][5].Value
	board.saa = sheet.Columns[1][6].Value

	return board
}

// Represents the editable fields on a Toastmasters agenda.
type agendaRoles struct {
	toastmaster       string
	ge                string
	timer             string
	ahCounter         string
	grammarian        string
	tableTopicsMaster string
	jokeMaster        string
	speakers          []speaker
	boardMembers      board
	futureWeeks       [][]string
}

// Factory method to create agenda roles from a google doc based on the date of the meeting.
func (agendaRoles) new(agendaDate string) (agendaRoles, error) {
	spreadsheets, err := getSheet()
	if err != nil {
		return agendaRoles{}, err
	}
	boardMembers := board{}.new(spreadsheets.boardSheet)

	agendaRoles := agendaRoles{}
	agendaRoles.boardMembers = boardMembers

	const speakerCellStart = 7
	const speakerCellEnd = 13
	rolesSheet := spreadsheets.meetingRoles
	for i := range rolesSheet.Columns {
		if rolesSheet.Columns[i][0].Value == agendaDate {
			agendaRoles.toastmaster = rolesSheet.Columns[i][1].Value
			agendaRoles.jokeMaster = rolesSheet.Columns[i][2].Value
			agendaRoles.ge = rolesSheet.Columns[i][3].Value
			agendaRoles.timer = rolesSheet.Columns[i][4].Value
			agendaRoles.ahCounter = rolesSheet.Columns[i][5].Value
			agendaRoles.grammarian = rolesSheet.Columns[i][6].Value

			for j := speakerCellStart; j <= speakerCellEnd; j += 2 {
				agendaRoles.speakers = append(agendaRoles.speakers, NewSpeaker(rolesSheet.Columns[i][j].Value,
					rolesSheet.Columns[i][j+1].Value))
			}

			agendaRoles.tableTopicsMaster = rolesSheet.Columns[i][16].Value
			agendaRoles.futureWeeks = getFutureWeeks(rolesSheet, i)
			break
		}
	}
	return agendaRoles, nil
}

//  Represents a speaker in a Toastmasters meeting.
type speaker struct {
	name string
	Speech
	evaluator string
}

// Factory method to create a speaker based on the spreadsheet speaker and evaluator.
func NewSpeaker(s string, eval string) speaker {
	name, manual, number := parseManualAndNumber(s)

	return speaker{
		name:      name,
		evaluator: eval,
		Speech:    NewSpeech(manual, number),
	}
}

// Helper method that returns the first name of a speaker.
func (s speaker) firstName() string {
	return strings.Split(s.name, " ")[0]
}

// Represents the spreadsheet tabs.
type googleDocsSheet struct {
	boardSheet   *spreadsheet.Sheet
	meetingRoles *spreadsheet.Sheet
}

//  GetSheet reads a Google Docs spreadsheet and returns a sheet with roles and another sheet with the board members.
func getSheet() (googleDocsSheet, error) {
	data, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		return googleDocsSheet{}, errors.New("cannot read client_secret.json")
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return googleDocsSheet{}, errors.New("problem with google.JWTConfigFromJSON(data, s.Scope)")
	}

	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	s, err := service.FetchSpreadsheet("1CBlORqCzL6YvyAUZTk8jezvhyuDzjjumghwGKk5VIK8")
	if err != nil {
		return googleDocsSheet{}, errors.New("cannot fetch spread sheet: ")
	}

	roles, err := s.SheetByIndex(0)
	if err != nil {
		return googleDocsSheet{}, errors.New("cannot read s by index 0")
	}

	board, err := s.SheetByIndex(1)
	if err != nil {
		return googleDocsSheet{}, errors.New("cannot read s by index 1")
	}

	return googleDocsSheet{boardSheet: board, meetingRoles: roles}, nil
}

// Find the speaker name, manual and number from a string that looks like "Ann Addicks\nCC #9".
func parseManualAndNumber(speaker string) (string, string, int) {
	re := regexp.MustCompile(`([a-zA-Z]+ [a-zA-Z]+)\n([a-zA-Z]+) #(\d{1,2})`)
	result := re.FindStringSubmatch(speaker)
	name := speaker
	manual := ""
	speechNum := 0

	if len(result) > 0 {
		name = result[1]
		manual = result[2]
		speechNum, _ = strconv.Atoi(result[3])
	}
	return name, manual, speechNum
}

// The number of weeks in the future to capture.
const futureWeeks = 4
const numberOfRoles = 17

// GetFutureWeeks finds the next several weeks after the current week based on the constant futureWeeks.
func getFutureWeeks(sheet *spreadsheet.Sheet, thisWeek int) [][]string {
	week := 0
	var nextSchedule = make([][]string, 0, futureWeeks)

	for i := thisWeek + 1; i < len(sheet.Columns) && week <= futureWeeks; i++ {
		nextWeek := make([]string, numberOfRoles)

		for j := 0; j < numberOfRoles; j++ {
			nextWeek[j] = sheet.Columns[i][j].Value
		}
		nextSchedule = append(nextSchedule, nextWeek)
		week++

	}
	return nextSchedule
}
