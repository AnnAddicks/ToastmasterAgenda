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

// Board members of a Toastmasters meeting.
type Board struct {
	President string
	VPE       string
	VPM       string
	VPPR      string
	Secretary string
	Treasurer string
	SAA       string
}

// NewBoard is a factory function using a spreadsheet to fill in Board members.
func NewBoard(sheet *spreadsheet.Sheet) *Board {
	return &Board{
		President: sheet.Columns[1][0].Value,
		VPE:       sheet.Columns[1][1].Value,
		VPM:       sheet.Columns[1][2].Value,
		VPPR:      sheet.Columns[1][3].Value,
		Secretary: sheet.Columns[1][4].Value,
		Treasurer: sheet.Columns[1][5].Value,
		SAA:       sheet.Columns[1][6].Value,
	}
}

// AgendaRoles contains the editable fields on a Toastmasters agenda.
type AgendaRoles struct {
	toastmaster       string
	ge                string
	timer             string
	ahCounter         string
	grammarian        string
	tableTopicsMaster string
	jokeMaster        string
	speakers          []*Speaker
	boardMembers      *Board
	futureWeeks       [][]string
}

// NewAgendaRoles is a factory function to create agenda roles from a google doc based on the date of the meeting.
func NewAgendaRoles(agendaDate string) (*AgendaRoles, error) {
	spreadsheets, err := fetchSheet()
	if err != nil {
		return &AgendaRoles{}, err
	}

	agendaRoles := &AgendaRoles{
		boardMembers: NewBoard(spreadsheets.boardSheet),
	}

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
			agendaRoles.futureWeeks = futureWeeks(rolesSheet, i)
			break
		}
	}
	return agendaRoles, nil
}

// A Speaker in a Toastmasters meeting.
type Speaker struct {
	name string
	*Speech
	evaluator string
}

// Helper method that returns the first name of a Speaker.
func (s *Speaker) firstName() string {
	return strings.Split(s.name, " ")[0]
}

// Find the Speaker name, manual and number from a string that looks like "Ann Addicks\nCC #9".
func parseManualAndNumber(speaker string) (string, string, int) {
	re := regexp.MustCompile(`([a-zA-Z]+ [a-zA-Z]+)\n([a-zA-Z]+) #(\d{1,2})`)
	result := re.FindStringSubmatch(speaker)
	name := speaker
	var manual string
	var speechNum int

	if len(result) > 0 {
		name = result[1]
		manual = result[2]
		speechNum, _ = strconv.Atoi(result[3])
	}
	return name, manual, speechNum
}

// NewSpeaker is a factory function to create a Speaker based on the spreadsheet Speaker and evaluator.
func NewSpeaker(s string, eval string) *Speaker {
	name, manual, number := parseManualAndNumber(s)

	return &Speaker{
		name:      name,
		evaluator: eval,
		Speech:    NewSpeech(manual, number),
	}
}

// Represents the spreadsheet tabs.
type googleDocsSheet struct {
	boardSheet   *spreadsheet.Sheet
	meetingRoles *spreadsheet.Sheet
}

//  GetSheet reads a Google Docs spreadsheet and returns a sheet with roles and another sheet with the Board members.
func fetchSheet() (googleDocsSheet, error) {
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

// FutureWeeks finds the next several weeks after the current week based on the constant futureWeeks.
func futureWeeks(sheet *spreadsheet.Sheet, thisWeek int) [][]string {
	// The number of weeks in the future to capture.
	const numOfWeeks = 4
	const numberOfRoles = 17

	var week int
	nextSchedule := make([][]string, 0, numOfWeeks)
	colLen := len(sheet.Columns)

	for i := thisWeek + 1; i < colLen && week <= numOfWeeks; i++ {
		nextWeek := make([]string, numberOfRoles)

		for j := 0; j < numberOfRoles; j++ {
			nextWeek[j] = sheet.Columns[i][j].Value
		}
		nextSchedule = append(nextSchedule, nextWeek)
		week++

	}
	return nextSchedule
}
