package main

import (
	"gopkg.in/Iwark/spreadsheet.v2"
	"testing"
)

func TestGetBoard(t *testing.T) {
	columns := make([][]spreadsheet.Cell, 2)
	columns[1] = make([]spreadsheet.Cell, 7)
	columns[1][0] = spreadsheet.Cell{Value: "president"}
	columns[1][1] = spreadsheet.Cell{Value: "vpe"}
	columns[1][2] = spreadsheet.Cell{Value: "vpm"}
	columns[1][3] = spreadsheet.Cell{Value: "vppr"}
	columns[1][4] = spreadsheet.Cell{Value: "secretary"}
	columns[1][5] = spreadsheet.Cell{Value: "treasurer"}
	columns[1][6] = spreadsheet.Cell{Value: "saa"}

	sheet := spreadsheet.Sheet{Columns: columns}
	board := getBoard(&sheet)

	if board.President != "president" {
		t.Error("Expected 'president', got ", board.President)
	}
	if board.Vpe != "vpe" {
		t.Error("Expected 'vpe', got ", board.Vpe)
	}
	if board.Vpm != "vpm" {
		t.Error("Expected 'vpm', got ", board.Vpm)
	}
	if board.Vppr != "vppr" {
		t.Error("Expected 'vppr', got ", board.Vppr)
	}
	if board.Secretary != "secretary" {
		t.Error("Expected 'secretary', got ", board.Secretary)
	}
	if board.Treasurer != "treasurer" {
		t.Error("Expected 'treasurer', got ", board.Treasurer)
	}
	if board.Saa != "saa" {
		t.Error("Expected 'saa', got ", board.Saa)
	}

}

func TestParseManualAndNumber(t *testing.T) {

	speech := "Ann Addicks\nCC #4 "

	name, manual, num := parseManualAndNumber(speech)

	if name != "Ann Addicks" {
		t.Error("Expected 'Ann Addicks', got ", name)
	}

	if manual != "CC" {
		t.Error("Expected 'CC', got ", manual)
	}

	if num != 4 {
		t.Error("Expected '4', got ", num)
	}
}
