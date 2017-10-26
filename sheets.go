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

type Board struct {
	President, Vpe, Vpm, Vppr, Secretary, Treasurer, Saa string
}

type AgendaRoles struct {
	Toastmaster, Ge, Timer, AhCounter, Grammarian string
	TableTopicsMaster, JokeMaster                 string
	Speakers                                      []Speaker
	BoardMembers                                  Board
	FutureWeeks                                   [][]string
}

type Speaker struct {
	Name      string
	Speech    Speech
	Evaluator string
}

func (s Speaker) firstName() string {
	return strings.Split(s.Name, " ")[0]
}

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

func getBoard(sheet *spreadsheet.Sheet) Board {
	board := Board{}
	board.President = sheet.Columns[1][0].Value
	board.Vpe = sheet.Columns[1][1].Value
	board.Vpm = sheet.Columns[1][2].Value
	board.Vppr = sheet.Columns[1][3].Value
	board.Secretary = sheet.Columns[1][4].Value
	board.Treasurer = sheet.Columns[1][5].Value
	board.Saa = sheet.Columns[1][6].Value

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

func populateSpeaker(s string, eval string) Speaker {
	name, manual, number := parseManualAndNumber(s)
	info := GetSpeech(manual, number)

	speaker := Speaker{}
	speaker.Name = name
	speaker.Evaluator = eval
	speaker.Speech = info
	return speaker
}

func GetRoles(agendaDate string) AgendaRoles {
	sheet, roles := getSheet()
	boardMembers := getBoard(roles)

	agendaRoles := AgendaRoles{}
	agendaRoles.BoardMembers = boardMembers

	for i := range sheet.Columns {
		if sheet.Columns[i][0].Value == agendaDate {
			agendaRoles.Toastmaster = sheet.Columns[i][1].Value
			agendaRoles.JokeMaster = sheet.Columns[i][2].Value
			agendaRoles.Ge = sheet.Columns[i][3].Value
			agendaRoles.Timer = sheet.Columns[i][4].Value
			agendaRoles.AhCounter = sheet.Columns[i][5].Value
			agendaRoles.Grammarian = sheet.Columns[i][6].Value

			for j := 7; j <= 13; j += 2 {
				agendaRoles.Speakers = append(agendaRoles.Speakers, populateSpeaker(sheet.Columns[i][j].Value, sheet.Columns[i][j+1].Value))
			}

			agendaRoles.TableTopicsMaster = sheet.Columns[i][16].Value
			agendaRoles.FutureWeeks = GetFutureWeeks(agendaDate, sheet)
			break
		}
	}
	return agendaRoles
}

func GetFutureWeeks(agendaDate string, sheet *spreadsheet.Sheet) [][]string {
	week := 0
	var nextSchedule = make([][]string, 0, 4)

	for i := 0; i < len(sheet.Columns) && week < 5; i++ {
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
