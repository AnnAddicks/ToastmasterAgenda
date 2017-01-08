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

	docx1 := r.Editable()
	date := "./" + AgendaDate() + ".docx"
	docx1.Replace("<<Date>>", date, -1)
	docx1.Replace("<<President>>", "Shelly Bowe", -1)
	docx1.Replace("<<VPE>>", "Ann Addicks", -1)
	docx1.WriteToFile(date)

	r.Close()
}
