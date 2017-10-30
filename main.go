package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nguyenthenguyen/docx"
)

func main() {
	now := time.Now()
	nt := nextTuesday(now)

	fmt.Println("Generating Agenda for", monthDayCommaYear(nt))
	createDoc(nt)
}

func createDoc(t time.Time) {
	r, err := docx.ReadDocxFile("./Agenda.docx")
	if err != nil {
		panic(err)
	}

	prettyPrintDate := monthDayCommaYear(t)
	dateWithPeriods := formatDate(t, delimiterPeriods)
	roles := agendaRoles{}.new(formatDate(t, delimiterSlashes))

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
	for i := range roles.speakers {
		speechOrder := i + 1
		soString := strconv.Itoa(speechOrder)
		speaker := roles.speakers[i]

		docx1.Replace("evaluator"+soString, speaker.evaluator, -1)
		docx1.Replace("speaker"+soString+"FirstLastName", speaker.name, -1)
		docx1.Replace("firstName"+soString, speaker.firstName(), -1)
		docx1.Replace("speaker"+soString+"Manual", speaker.speech.manualName, -1)
		docx1.Replace("speaker"+soString+"Speech", speaker.speech.info(), -1)

		// Replace speech times for the second through fourth speaker based on last max speech time plus one.
		if speechOrder == 1 {
			curTime := time.Date(2017, time.January, 1, 7, 14, 0, 0, time.UTC)
			nextTime, _ = addMinutes(curTime, 0)
			pastSpeechTime = speaker.speech.max + 1

		} else {
			nextTime, printString = addMinutes(nextTime, pastSpeechTime)
			docx1.Replace("e"+soString+"t"+soString, printString, 1)

			nextTime, printString = addMinutes(nextTime, +1)
			docx1.Replace("s"+soString+"t"+soString, printString, 1)
			pastSpeechTime = speaker.speech.max + 1
		}
	}
	docx1.Replace("tTMaster", roles.tableTopicsMaster, -1)
	_, printString = addMinutes(nextTime, pastSpeechTime)
	docx1.Replace("ttmt", printString, 1)

	// Replace the next several weeks on the agenda.
	for i := range roles.futureWeeks {
		nextWeek := roles.futureWeeks[i]

		for j := range nextWeek {
			docx1.Replace("w"+strconv.Itoa(i)+"_"+strconv.Itoa(j), nextWeek[j], 1)
		}
	}

	docx1.WriteToFile(fileName)
	r.Close()
}
