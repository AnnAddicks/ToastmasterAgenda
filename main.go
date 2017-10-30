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
	roles := GetRoles(formatDate(t, delimiterSlashes))

	docx1 := r.Editable()
	fileName := "./" + dateWithPeriods + ".docx"

	docx1.ReplaceHeader("Date", prettyPrintDate)
	docx1.Replace("president", roles.BoardMembers.President, -1)
	docx1.Replace("vpe", roles.BoardMembers.Vpe, -1)
	docx1.Replace("vpm", roles.BoardMembers.Vpm, -1)
	docx1.Replace("vppr", roles.BoardMembers.Vppr, -1)
	docx1.Replace("secretary", roles.BoardMembers.Secretary, -1)
	docx1.Replace("treasurer", roles.BoardMembers.Treasurer, -1)
	docx1.Replace("saa", roles.BoardMembers.Saa, -1)
	docx1.Replace("jokeMaster", roles.JokeMaster, -1)
	docx1.Replace("toastmasterOfDay", roles.Toastmaster, -1)
	docx1.Replace("generalEval", roles.Ge, -1)
	docx1.Replace("timer", roles.Timer, -1)
	docx1.Replace("ah-counter", roles.AhCounter, -1)
	docx1.Replace("grammarian", roles.Grammarian, -1)

	var nextTime time.Time
	var pastSpeechTime int
	var printString string
	for i := range roles.Speakers {
		speechOrder := i + 1
		speaker := roles.Speakers[i]

		docx1.Replace("evaluator"+strconv.Itoa(speechOrder), speaker.Evaluator, -1)
		docx1.Replace("speaker"+strconv.Itoa(speechOrder)+"FirstLastName", speaker.Name, -1)
		docx1.Replace("firstName"+strconv.Itoa(speechOrder), speaker.firstName(), -1)
		docx1.Replace("speaker"+strconv.Itoa(speechOrder)+"Manual", speaker.Speech.manualName, -1)
		docx1.Replace("speaker"+strconv.Itoa(speechOrder)+"Speech", speaker.Speech.info(), -1)

		// Replace speech times for the second through fourth speaker based on last max speech time plus one.
		if speechOrder == 1 {
			curTime := time.Date(2017, time.January, 1, 7, 14, 0, 0, time.UTC)
			nextTime, _ = addMinutes(curTime, 0)
			pastSpeechTime = speaker.Speech.max + 1

		} else {
			nextTime, printString = addMinutes(nextTime, pastSpeechTime)
			docx1.Replace("e"+strconv.Itoa(speechOrder)+"t"+strconv.Itoa(speechOrder), printString, 1)

			nextTime, printString = addMinutes(nextTime, +1)
			docx1.Replace("s"+strconv.Itoa(speechOrder)+"t"+strconv.Itoa(speechOrder), printString, 1)
			pastSpeechTime = speaker.Speech.max + 1
		}
	}
	docx1.Replace("tTMaster", roles.TableTopicsMaster, -1)
	_, printString = addMinutes(nextTime, pastSpeechTime)
	docx1.Replace("ttmt", printString, 1)

	// Replace the next several weeks on the agenda.
	for i := range roles.FutureWeeks {
		nextWeek := roles.FutureWeeks[i]

		for j := range nextWeek {
			docx1.Replace("w"+strconv.Itoa(i)+"_"+strconv.Itoa(j), nextWeek[j], 1)
		}
	}

	docx1.WriteToFile(fileName)
	r.Close()
}
