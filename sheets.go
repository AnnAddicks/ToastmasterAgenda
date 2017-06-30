package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

func parseManualAndNumber(speaker string) (string, int) {
	           				// `(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`
	re := regexp.MustCompile(`(?P<manual>\n[a-zA-Z]+)\s(?P<number>#\d{1,2})`)

	result_slice := re.FindAllStringSubmatch(speaker, -1)
	fmt.Printf("%v", result_slice)

	fmt.Println("first slice: ", result_slice[0])
	

	manual := "CC"
	speechNum := 1

	return manual, speechNum
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

			agendaRoles.speaker1 = sheet.Columns[i][7].Value
			agendaRoles.speaker1FirstName = strings.Split(agendaRoles.speaker1, " ")[0]
			agendaRoles.eval1 = sheet.Columns[i][8].Value

			manual, number := parseManualAndNumber(agendaRoles.speaker1)
			speech := GetSpeech(manual, number)
			agendaRoles.speaker1Manual = speech.manualName
			agendaRoles.speaker1Speech = speech.name //add the times too!

			agendaRoles.speaker2 = sheet.Columns[i][9].Value
			agendaRoles.speaker2FirstName = strings.Split(agendaRoles.speaker2, " ")[0]
			agendaRoles.eval2 = sheet.Columns[i][10].Value

			agendaRoles.speaker3 = sheet.Columns[i][11].Value
			agendaRoles.speaker3FirstName = strings.Split(agendaRoles.speaker3, " ")[0]
			agendaRoles.eval3 = sheet.Columns[i][12].Value

			agendaRoles.speaker4 = sheet.Columns[i][13].Value
			agendaRoles.speaker4FirstName = strings.Split(agendaRoles.speaker1, " ")[0]
			agendaRoles.eval4 = sheet.Columns[i][14].Value

			agendaRoles.tableTopicsMaster = sheet.Columns[i][18].Value
			break
		}
	}
	return agendaRoles
}

func GetFutureWeeks(agendaDate string, sheet *spreadsheet.Sheet) {
	week := 0

	for i := range sheet.Columns {

		if week > 0 {

		}

		if sheet.Columns[i][0].Value == agendaDate {
			week = 1
		}

	}
}
