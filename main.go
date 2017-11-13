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

// Creates a word document with the Name DD.MM.YYYY.docx based on Agenda.docx.
func createDoc(t time.Time) error {
	r, err := docx.ReadDocxFile("./Agenda.docx")
	defer r.Close()

	if err != nil {
		return err
	}

	d := r.Editable()
	if err := replaceFields(d, t); err != nil {
		return err
	}

	d.WriteToFile("./" + formatDate(t, delimiterPeriods) + ".docx")
	return nil
}

// ReplaceSpeakers replaces the speech time, speaker, speech, Evaluator time, and Evaluator.
func replaceSpeakers(d *docx.Docx, s []*Speaker) {
	// Time for Speech evaluation goals starts at 7:13 pm.
	curTime := time.Date(2017, time.January, 1, 7, 13, 0, 0, time.UTC)
	nextTime, _ := addMinutes(curTime, 0)
	var pastSpeechTime int
	var printString string
	for i := range s {
		speechOrder := i + 1
		soString := strconv.Itoa(speechOrder)
		speaker := s[i]

		d.Replace("evaluator"+soString, speaker.Evaluator, -1)
		d.Replace("speaker"+soString+"FirstLastName", speaker.Name, -1)
		d.Replace("firstName"+soString, speaker.firstName(), -1)
		d.Replace("speaker"+soString+"Manual", speaker.Speech.ManualName, -1)
		d.Replace("speaker"+soString+"Speech", speaker.Speech.info(), -1)

		// Replace Speech times for Speaker and Evaluator based on last Max Speech time plus one.
		nextTime, printString = addMinutes(nextTime, pastSpeechTime)
		d.Replace("e"+soString+"t"+soString, printString, 1)

		nextTime, printString = addMinutes(nextTime, +1)
		d.Replace("s"+soString+"t"+soString, printString, 1)
		pastSpeechTime = speaker.Speech.Max + 1
	}

	_, ttTime := addMinutes(nextTime, pastSpeechTime)
	d.Replace("ttmt", ttTime, 1)
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
	d.Replace("president", roles.BoardMembers.President, -1)
	d.Replace("vpe", roles.BoardMembers.VPE, -1)
	d.Replace("vpm", roles.BoardMembers.VPM, -1)
	d.Replace("vppr", roles.BoardMembers.VPPR, -1)
	d.Replace("secretary", roles.BoardMembers.Secretary, -1)
	d.Replace("treasurer", roles.BoardMembers.Treasurer, -1)
	d.Replace("saa", roles.BoardMembers.SAA, -1)
	d.Replace("jokeMaster", roles.JokeMaster, -1)
	d.Replace("toastmasterOfDay", roles.Toastmaster, -1)
	d.Replace("generalEval", roles.GE, -1)
	d.Replace("timer", roles.Timer, -1)
	d.Replace("ah-counter", roles.AhCounter, -1)
	d.Replace("grammarian", roles.Grammarian, -1)
	d.Replace("tTMaster", roles.TableTopicsMaster, -1)

	replaceSpeakers(d, roles.Speakers)
	replaceFutureWeeks(d, roles.FutureWeeks)
	return nil
}
