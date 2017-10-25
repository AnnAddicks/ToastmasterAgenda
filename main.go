package main

import (
	"fmt"
	"github.com/annaddicks/docx"
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

	docx1.ReplaceHeader("Date", prettyPrintDate)
	docx1.Replace("president", roles.boardMembers.president, -1)
	docx1.Replace("vpe", roles.boardMembers.vpe, -1)
	docx1.Replace("vpm", roles.boardMembers.vpm, -1)
	docx1.Replace("vppr", roles.boardMembers.vppr, -1)
	docx1.Replace("secretary", roles.boardMembers.secretary, -1)
	docx1.Replace("treasurer", roles.boardMembers.treasurer, -1)
	docx1.Replace("saa", roles.boardMembers.saa, -1)
	docx1.Replace("jokeMaster", roles.jokeMaster, -1)
	docx1.Replace("toastmasterOfDay", roles.toastmaster, -1)
	docx1.Replace("generalEval", roles.ge, -1)
	docx1.Replace("timer", roles.timer, -1)
	docx1.Replace("ah-counter", roles.ahCounter, -1)
	docx1.Replace("grammarian", roles.grammarian, -1)

	var nextTime time.Time
	var pastSpeechTime int
	var printString string
	for i := 0; i < 4; i++ {
		speachOrder := i + 1
		speaker := roles.speakers[i]

		docx1.Replace("evaluator"+strconv.Itoa(speachOrder), speaker.Evaluator, -1)
		docx1.Replace("speaker"+strconv.Itoa(speachOrder)+"FirstLastName", speaker.Name, -1)
		docx1.Replace("firstName"+strconv.Itoa(speachOrder), speaker.firstName(), -1)
		docx1.Replace("speaker"+strconv.Itoa(speachOrder)+"Manual", speaker.Speech.manualName, -1)
		docx1.Replace("speaker"+strconv.Itoa(speachOrder)+"Speech", speaker.Speech.name, -1)

		//Replace speech times for the second - fourth speaker
		if speachOrder == 1 {
			curTime := time.Date(2017, time.January, 1, 7, 14, 0, 0, time.UTC)
			nextTime, _ = prettyPrintTime(curTime, 0)
			pastSpeechTime = speaker.Speech.max + 1

		} else {
			nextTime, printString = prettyPrintTime(nextTime, pastSpeechTime)
			docx1.Replace("e"+strconv.Itoa(speachOrder)+"t"+strconv.Itoa(speachOrder), printString, 1)

			nextTime, printString = prettyPrintTime(nextTime, +1)
			docx1.Replace("s"+strconv.Itoa(speachOrder)+"t"+strconv.Itoa(speachOrder), printString, 1)
			pastSpeechTime = speaker.Speech.max + 1
		}
	}
	docx1.Replace("tTMaster", roles.tableTopicsMaster, -1)
	_, printString = prettyPrintTime(nextTime, pastSpeechTime)
	docx1.Replace("ttmt", printString, 1)

	//Replace the next 4 weeks
	for i := range roles.futureWeeks {
		nextWeek := roles.futureWeeks[i]

		for j := 0; j < 17; j++ {
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
	createDoc(t)
}
