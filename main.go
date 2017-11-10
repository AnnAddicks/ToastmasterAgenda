package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nguyenthenguyen/docx"
)

// Go forth and create an agenda.
func main() {
	nt := nextTuesday(time.Now())
	fmt.Println("Generating Agenda for", monthDayCommaYear(nt))

	if err := createDoc(nt); err != nil {
		panic(err)
	}
}

// Creates a word document with the name DD.MM.YYYY.docx based on Agenda.docx.
func createDoc(t time.Time) error {
	r, err := docx.ReadDocxFile("./Agenda.docx")
	if err != nil {
		return err
	}

	d := r.Editable()
	if err := replaceFields(d, t); err != nil {
		return err
	}

	d.WriteToFile("./" + formatDate(t, delimiterPeriods) + ".docx")
	r.Close()

	return nil
}

// ReplaceSpeakers replaces the speech time, speaker, speech, evaluator time, and evaluator.
func replaceSpeakers(d *docx.Docx, s []*Speaker) string {
	// Time for Speech evaluation goals starts at 7:13 pm.
	curTime := time.Date(2017, time.January, 1, 7, 13, 0, 0, time.UTC)
	nextTime, _ := addMinutes(curTime, 0)
	var pastSpeechTime int
	var printString string
	for i := range s {
		speechOrder := i + 1
		soString := strconv.Itoa(speechOrder)
		speaker := s[i]

		d.Replace("evaluator"+soString, speaker.evaluator, -1)
		d.Replace("speaker"+soString+"FirstLastName", speaker.name, -1)
		d.Replace("firstName"+soString, speaker.firstName(), -1)
		d.Replace("speaker"+soString+"Manual", speaker.Speech.manualName, -1)
		d.Replace("speaker"+soString+"Speech", speaker.Speech.info(), -1)

		// Replace Speech times for Speaker and evaluator based on last max Speech time plus one.
		nextTime, printString = addMinutes(nextTime, pastSpeechTime)
		d.Replace("e"+soString+"t"+soString, printString, 1)

		nextTime, printString = addMinutes(nextTime, +1)
		d.Replace("s"+soString+"t"+soString, printString, 1)
		pastSpeechTime = speaker.Speech.max + 1
	}

	_, lastSpeechTime := addMinutes(nextTime, pastSpeechTime)
	return lastSpeechTime
}

// ReplaceFutureWeeks replaces the future schedule on the agenda.
func replaceFutureWeeks(d *docx.Docx, fw [][]string) {
	for i := range fw {
		nextWeek := fw[i]

		for j := range nextWeek {
			d.Replace("w"+strconv.Itoa(i)+"_"+strconv.Itoa(j), nextWeek[j], 1)
		}
	}
}

// ReplaceFields handles the replacement of all the placeholders in the agenda.
func replaceFields(d *docx.Docx, t time.Time) error {
	roles, err := NewAgendaRoles(formatDate(t, delimiterSlashes))

	if err != nil {
		return err
	}

	d.ReplaceHeader("Date", monthDayCommaYear(t))
	d.Replace("president", roles.boardMembers.president, -1)
	d.Replace("vpe", roles.boardMembers.vpe, -1)
	d.Replace("vpm", roles.boardMembers.vpm, -1)
	d.Replace("vppr", roles.boardMembers.vppr, -1)
	d.Replace("secretary", roles.boardMembers.secretary, -1)
	d.Replace("treasurer", roles.boardMembers.treasurer, -1)
	d.Replace("saa", roles.boardMembers.saa, -1)
	d.Replace("jokeMaster", roles.jokeMaster, -1)
	d.Replace("toastmasterOfDay", roles.toastmaster, -1)
	d.Replace("generalEval", roles.ge, -1)
	d.Replace("timer", roles.timer, -1)
	d.Replace("ah-counter", roles.ahCounter, -1)
	d.Replace("grammarian", roles.grammarian, -1)
	d.Replace("tTMaster", roles.tableTopicsMaster, -1)

	ttTime := replaceSpeakers(d, roles.speakers)
	d.Replace("ttmt", ttTime, 1)
	replaceFutureWeeks(d, roles.futureWeeks)
	return nil
}
