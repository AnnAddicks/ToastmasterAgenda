package main

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
)

func main() {
	fmt.Println("Hello World")

	r, err := docx.ReadDocxFile("./Agenda.docx")
	if err != nil {
		panic(err)
	}

	roleDate := AgendaDayMonthYear()
	roles := GetRoles(roleDate)

	docx1 := r.Editable()
	date := "./" + AgendaDate() + ".docx"

	docx1.Replace("Date", date, -1)
	docx1.Replace("president", roles.boardMembers.president, -1)
	docx1.Replace("vpe", roles.boardMembers.vpe, -1)
        docx1.Replace("vpm", roles.boardMembers.vpm, -1)
	docx1.Replace("vppr", roles.boardMembers.vppr, -1)
	docx1.Replace("secretary", roles.boardMembers.secretary, -1)
	docx1.Replace("treasurer", roles.boardMembers.treasurer, -1)
	docx1.Replace("saa", roles.boardMembers.saa, -1)
	docx1.Replace("toastmaster", roles.toastmaster, -1)
	docx1.Replace("GE", roles.ge, -1)
	docx1.Replace("Timer", roles.timer, -1)
	docx1.Replace("Ah-Counter", roles.ahCounter, -1)
	docx1.Replace("Grammarian", roles.grammarian, -1)
	docx1.Replace("Evaluator1", roles.eval1, -1)
	docx1.Replace("Speaker1", roles.speaker1, -1)
	docx1.Replace("FirstName1", roles.speaker1FirstName, -1)
	docx1.Replace("Evaluator2", roles.eval2, -1)
	docx1.Replace("<<Speaker2>>", roles.speaker2, -1)
	docx1.Replace("<<FirstName2>>", roles.speaker2FirstName, -1)
	docx1.Replace("<<Evaluator3>>", roles.eval3, -1)
	docx1.Replace("<<Speaker3>>", roles.speaker3, -1)
	docx1.Replace("<<FirstName3>>", roles.speaker3FirstName, -1)
	docx1.Replace("<<Evaluator4>>", roles.eval4, -1)
	docx1.Replace("<<Speaker4>>", roles.speaker4, -1)
	docx1.Replace("<<FirstName4>>", roles.speaker4FirstName, -1)
	docx1.Replace("TTMaster", roles.tableTopicsMaster, -1)

	docx1.WriteToFile(date)
	r.Close()
}
