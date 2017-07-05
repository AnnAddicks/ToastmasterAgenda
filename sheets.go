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
	president, vpe, vpm, vppr, secretary, treasurer, saa string
}
type AgendaRoles struct {
	toastmaster, ge, timer, ahCounter, grammarian, eval1, speaker1, speaker1FirstName, speaker1Manual, speaker1Speech      string
	eval2, speaker2, speaker2FirstName, speaker2Manual, speaker2Speech, eval3, speaker3, speaker3FirstName, speaker3Manual string
	speaker3Speech, eval4, speaker4, speaker4FirstName, speaker4Manual, speaker4Speech, tableTopicsMaster                  string
	boardMembers                                                                                                           Board
	futureWeeks                                                                                                            [4][16]string
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
	re := regexp.MustCompile(`([a-zA-Z]+ [a-zA-Z]+)\n(?P<manual>[a-zA-Z]+) #(?P<number>\d{1,2})`)
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

func GetRoles(agendaDate string) AgendaRoles {
	sheet, roles := getSheet()
	boardMembers := getBoard(roles)

	agendaRoles := AgendaRoles{}
	agendaRoles.boardMembers = boardMembers

	for i := range sheet.Columns {
		if sheet.Columns[i][0].Value == agendaDate {
			agendaRoles.toastmaster = sheet.Columns[i][1].Value
			agendaRoles.ge = sheet.Columns[i][2].Value
			agendaRoles.timer = sheet.Columns[i][3].Value
			agendaRoles.ahCounter = sheet.Columns[i][4].Value
			agendaRoles.grammarian = sheet.Columns[i][5].Value

			//MAJOR CPD, pullout a method, possibly with a nested
			name, manual, number := parseManualAndNumber(sheet.Columns[i][7].Value)
			speech := GetSpeech(manual, number)
			agendaRoles.speaker1 = name
			agendaRoles.speaker1FirstName = strings.Split(agendaRoles.speaker1, " ")[0]
			agendaRoles.eval1 = sheet.Columns[i][8].Value
			agendaRoles.speaker1Manual = speech.manualName
			agendaRoles.speaker1Speech = speech.name

			name, manual, number = parseManualAndNumber(sheet.Columns[i][9].Value)
			speech = GetSpeech(manual, number)
			agendaRoles.speaker2 = name
			agendaRoles.speaker2FirstName = strings.Split(agendaRoles.speaker2, " ")[0]
			agendaRoles.eval2 = sheet.Columns[i][10].Value
			agendaRoles.speaker2Manual = speech.manualName
			agendaRoles.speaker2Speech = speech.name

			name, manual, number = parseManualAndNumber(sheet.Columns[i][11].Value)
			speech = GetSpeech(manual, number)
			agendaRoles.speaker3 = name
			agendaRoles.speaker3FirstName = strings.Split(agendaRoles.speaker3, " ")[0]
			agendaRoles.eval3 = sheet.Columns[i][12].Value
			agendaRoles.speaker3Manual = speech.manualName
			agendaRoles.speaker3Speech = speech.name

			name, manual, number = parseManualAndNumber(sheet.Columns[i][13].Value)
			speech = GetSpeech(manual, number)
			agendaRoles.speaker4 = name
			agendaRoles.speaker4FirstName = strings.Split(agendaRoles.speaker4, " ")[0]
			agendaRoles.eval4 = sheet.Columns[i][14].Value
			agendaRoles.speaker4Manual = speech.manualName
			agendaRoles.speaker4Speech = speech.name

			agendaRoles.tableTopicsMaster = sheet.Columns[i][18].Value

			agendaRoles.futureWeeks = GetFutureWeeks(agendaDate, sheet)
			break
		}
	}
	return agendaRoles
}

func GetFutureWeeks(agendaDate string, sheet *spreadsheet.Sheet) [4][16]string {
	week := 0
	nextSchedule := [4][16]string{}

	for i := range sheet.Columns {

		if week == 0 {
			if sheet.Columns[i][0].Value == agendaDate {
				week = 1
			}

		} else if week == 5 {
			break
		} else {
			nextWeek := [16]string{}
			nextWeek[0] = sheet.Columns[i][0].Value
			nextWeek[1] = sheet.Columns[i][1].Value
			nextWeek[2] = sheet.Columns[i][2].Value
			nextWeek[3] = sheet.Columns[i][3].Value
			nextWeek[4] = sheet.Columns[i][4].Value
			nextWeek[5] = sheet.Columns[i][5].Value
			nextWeek[6] = sheet.Columns[i][7].Value
			nextWeek[7] = sheet.Columns[i][8].Value
			nextWeek[8] = sheet.Columns[i][9].Value
			nextWeek[9] = sheet.Columns[i][10].Value
			nextWeek[10] = sheet.Columns[i][11].Value
			nextWeek[11] = sheet.Columns[i][12].Value
			nextWeek[12] = sheet.Columns[i][13].Value
			nextWeek[13] = sheet.Columns[i][14].Value
			nextWeek[14] = sheet.Columns[i][16].Value
			nextWeek[15] = sheet.Columns[i][18].Value

			nextSchedule[week-1] = nextWeek
			week++

		}

	}

	return nextSchedule
}
