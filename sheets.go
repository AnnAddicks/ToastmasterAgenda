package main

import (
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
	president, vpe, vpm, vppr, secretary, treasurer, saa string
}

// Represents the editable fields on a Toastmasters agenda.
type agendaRoles struct {
	toastmaster, ge, timer, ahCounter, grammarian string
	tableTopicsMaster, jokeMaster                 string
	speakers                                      []speaker
	boardMembers                                  board
	futureWeeks                                   [][]string
}

//  Represents a speaker in a Toastmasters meeting.
type speaker struct {
	name      string
	speech
	evaluator string
}

// Helper method that returns the first name of a speaker.
func (s speaker) firstName() string {
	return strings.Split(s.name, " ")[0]
}

//  GetSheet reads a Google Docs spreadsheet and returns a sheet with roles and another sheet with the board members.
func getSheet() (*spreadsheet.Sheet, *spreadsheet.Sheet) {
	data, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		panic("cannot read client_secret.json")
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		panic("problem with google.JWTConfigFromJSON(data, spreadsheet.Scope)")
	}

	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet("1CBlORqCzL6YvyAUZTk8jezvhyuDzjjumghwGKk5VIK8")
	if err != nil {
		panic("cannot fetch spread sheet: ")
	}

	roles, err := spreadsheet.SheetByIndex(0)
	if err != nil {
		panic("Cannot read spreadsheet by index 0")
	}

	board, err := spreadsheet.SheetByIndex(1)
	if err != nil {
		panic("Cannot read spreadsheet by index 1")
	}

	return roles, board
}

func getBoard(sheet *spreadsheet.Sheet) board {
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

func populateSpeaker(s string, eval string) speaker {
	name, manual, number := parseManualAndNumber(s)
	info := speech{}.new(manual, number)

	speaker := speaker{}
	speaker.name = name
	speaker.evaluator = eval
	speaker.speech = info

	return speaker
}

func getRoles(agendaDate string) agendaRoles {
	sheet, roles := getSheet()
	boardMembers := getBoard(roles)

	agendaRoles := agendaRoles{}
	agendaRoles.boardMembers = boardMembers

	for i := range sheet.Columns {
		if sheet.Columns[i][0].Value == agendaDate {
			agendaRoles.toastmaster = sheet.Columns[i][1].Value
			agendaRoles.jokeMaster = sheet.Columns[i][2].Value
			agendaRoles.ge = sheet.Columns[i][3].Value
			agendaRoles.timer = sheet.Columns[i][4].Value
			agendaRoles.ahCounter = sheet.Columns[i][5].Value
			agendaRoles.grammarian = sheet.Columns[i][6].Value

			for j := 7; j <= 13; j += 2 {
				agendaRoles.speakers = append(agendaRoles.speakers, populateSpeaker(sheet.Columns[i][j].Value, sheet.Columns[i][j+1].Value))
			}

			agendaRoles.tableTopicsMaster = sheet.Columns[i][16].Value
			agendaRoles.futureWeeks = getFutureWeeks(agendaDate, sheet)
			break
		}
	}
	return agendaRoles
}

// The number of weeks in the future to capture.
const futureWeeks = 4

// GetFutureWeeks finds the next several weeks after the current week based on the constant futureWeeks.
func getFutureWeeks(agendaDate string, sheet *spreadsheet.Sheet) [][]string {
	week := 0
	var nextSchedule = make([][]string, 0, futureWeeks)

	for i := 0; i < len(sheet.Columns) && week <= futureWeeks; i++ {
		if week == 0 {
			if sheet.Columns[i][0].Value == agendaDate {
				week = 1
			}
		} else {
			nextWeek := make([]string, 17)

			for j := 0; j < 17; j++ {
				nextWeek[j] = sheet.Columns[i][j].Value
			}
			nextSchedule = append(nextSchedule, nextWeek)
			week++
		}
	}
	return nextSchedule
}
