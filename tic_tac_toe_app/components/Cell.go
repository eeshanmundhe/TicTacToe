package components

import (
	"errors"
)

//Cell takes Mark of string type
type Cell struct {
	Mark string
}

const (
	//NoMark stores the Mark for unused cell
	NoMark = "-"

	//XMark is used to assign X to the selected Cell
	XMark = "X"

	//OMark is used to assign O to the selected Cell
	OMark = "O"
)

//NewCell returns the Mark from Cell structmark
func NewCell() *Cell {
	return &Cell{Mark: NoMark}
}

//SetMark returns error if cell is already marked, else it marks the cell
func (cell *Cell) SetMark(mark string) error {
	if cell.Mark == NoMark {
		cell.Mark = mark
		return nil
	}
	return errors.New("Cell is already marked")
}

//GetMark returns the Mark of the cell
func (cell *Cell) GetMark() string {
	return cell.Mark
}
