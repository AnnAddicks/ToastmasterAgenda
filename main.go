package main

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
	"strconv"
	"time"
)

func createDoc(t time.Time) {
	r, err := docx.ReadDocxFile("./Agenda.docx")
	if err != nil {
		panic(err)
	}

	prettyPrintDate := AgendaMonthDayYear(t)
	dateWithPeriods := AgendaDate(t)
	roles := GetRoles(DateWithSlashes(t))

	docx1 := r.Editable()
	fileName := "./" + dateWithPeriods + ".docx"

	docx1.Replace("Date", prettyPrintDate, -1)
	docx1.Replace("president", roles.boardMembers.president, -1)
	docx1.Replace("vpe", roles.boardMembers.vpe, -1)
	docx1.Replace("vpm", roles.boardMembers.vpm, -1)
	docx1.Replace("vppr", roles.boardMembers.vppr, -1)
	docx1.Replace("secretary", roles.boardMembers.secretary, -1)
	docx1.Replace("treasurer", roles.boardMembers.treasurer, -1)
	docx1.Replace("saa", roles.boardMembers.saa, -1)
	docx1.Replace("toastmaster", roles.toastmaster, -1)
	docx1.Replace("generalEval", roles.ge, -1)
	docx1.Replace("timer", roles.timer, -1)
	docx1.Replace("ah-counter", roles.ahCounter, -1)
	docx1.Replace("grammarian", roles.grammarian, -1)
	docx1.Replace("evaluator1", roles.eval1, -1)
	docx1.Replace("speaker1FirstLastName", roles.speaker1, -1)
	docx1.Replace("firstName1", roles.speaker1FirstName, -1)
	docx1.Replace("speaker1Manual", roles.speech1.manualName, -1)
	docx1.Replace("speaker1Speech", roles.speech1.name, -1)
	docx1.Replace("evaluator2", roles.eval2, -1)
	docx1.Replace("speaker2FirstLastName", roles.speaker2, -1)
	docx1.Replace("firstName2", roles.speaker2FirstName, -1)
	docx1.Replace("speaker2Manual", roles.speech2.manualName, -1)
	docx1.Replace("speaker2Speech", roles.speech2.name, -1)
	docx1.Replace("evaluator3", roles.eval3, -1)
	docx1.Replace("speaker3FirstLastName", roles.speaker3, -1)
	docx1.Replace("firstName3", roles.speaker3FirstName, -1)
	docx1.Replace("speaker3Manual", roles.speech3.manualName, -1)
	docx1.Replace("speaker3Speech", roles.speech3.name, -1)
	docx1.Replace("evaluator4", roles.eval4, -1)
	docx1.Replace("speaker4FirstLastName", roles.speaker4, -1)
	docx1.Replace("firstName4", roles.speaker4FirstName, -1)
	docx1.Replace("speaker4Manual", roles.speech4.manualName, -1)
	docx1.Replace("speaker4Speech", roles.speech4.name, -1)
	docx1.Replace("tTMaster", roles.tableTopicsMaster, -1)

	//Replace speech & evaluator times
	curTime := time.Date(2017, time.January, 1, 7, 13, 0, 0, time.UTC)
	nextTime, printString := prettyPrintTime(curTime, roles.speech1.max+1)
	docx1.Replace("e2t2", printString, 1)

	nextTime, printString = prettyPrintTime(nextTime, +1)
	docx1.Replace("s2t2", printString, 1)

	nextTime, printString = prettyPrintTime(nextTime, roles.speech2.max+1)
	docx1.Replace("e3t3", printString, 1)

	nextTime, printString = prettyPrintTime(nextTime, 1)
	docx1.Replace("s3t3", printString, 1)

	nextTime, printString = prettyPrintTime(nextTime, roles.speech3.max+1)
	docx1.Replace("e4t4", printString, 1)

	nextTime, printString = prettyPrintTime(nextTime, 1)
	docx1.Replace("s4t4", printString, 1)

	_, printString = prettyPrintTime(nextTime, roles.speech4.max+1)
	docx1.Replace("ttmt", printString, 1)

	//Replace the next 4 weeks
	for i := range roles.futureWeeks {
		nextWeek := roles.futureWeeks[i]

		for j := 0; j < 16; j++ {
			docx1.Replace("w"+strconv.Itoa(i)+"_"+strconv.Itoa(j), nextWeek[j], 1)
		}
	}

	docx1.WriteToFile(fileName)
	r.Close()
}

func main() {
	d := time.Now()
	t := getNextTuesday(d)

	fmt.Println("Generating Agenda for", AgendaMonthDayYear(t))
	//fmt.Println("Press <Enter> to generate an agenda for", AgendaMonthDayYear(t))
	//fmt.Println("or type a new date with the format 'MM/DD/YYYY' and press <Enter>.")

	/*
	 *  Get Input
	 *  func TrimSpace(s string) string	clean up input
	 *  Verify correct string or enter
	 *
	 */

	createDoc(t)
}
