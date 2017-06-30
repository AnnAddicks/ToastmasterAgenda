package main

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
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
	docx1.Replace("ge", roles.ge, -1)
	docx1.Replace("timer", roles.timer, -1)
	docx1.Replace("ah-counter", roles.ahCounter, -1)
	docx1.Replace("grammarian", roles.grammarian, -1)
	docx1.Replace("evaluator1", roles.eval1, -1)
	docx1.Replace("speaker1", roles.speaker1, -1)
	docx1.Replace("firstName1", roles.speaker1FirstName, -1)
	docx1.Replace("Speaker1Manual", roles.speaker1Manual, -1)
	docx1.Replace("speaker1Speech", roles.speaker1Speech, -1)
	docx1.Replace("evaluator2", roles.eval2, -1)
	docx1.Replace("speaker2", roles.speaker2, -1)
	docx1.Replace("firstName2", roles.speaker2FirstName, -1)
	docx1.Replace("speaker2Manual", roles.speaker2Manual, -1)
	docx1.Replace("speaker2Speech", roles.speaker2Speech, -1)
	docx1.Replace("evaluator3", roles.eval3, -1)
	docx1.Replace("speaker3", roles.speaker3, -1)
	docx1.Replace("firstName3", roles.speaker3FirstName, -1)
	docx1.Replace("speaker3Manual", roles.speaker3Manual, -1)
	docx1.Replace("speaker3Speech", roles.speaker3Speech, -1)
	docx1.Replace("evaluator4", roles.eval4, -1)
	docx1.Replace("speaker4", roles.speaker4, -1)
	docx1.Replace("firstName4", roles.speaker4FirstName, -1)
	docx1.Replace("speaker4Manual", roles.speaker4Manual, -1)
	docx1.Replace("speaker4Speech", roles.speaker4Speech, -1)
	docx1.Replace("tTMaster", roles.tableTopicsMaster, -1)

	docx1.WriteToFile(fileName)
	r.Close()
}

func main() {
	d := time.Now()
	t := getNextTuesday(d)

	fmt.Println("Press <Enter> to generate an agenda for", AgendaMonthDayYear(t))
	fmt.Println("or type a new date with the format 'MM/DD/YYYY' and press <Enter>.")

	/*
	 *  Get Input
	 *  func TrimSpace(s string) string	clean up input
	 *  Verify correct string or enter
	 *
	 */

	createDoc(t)
}
