package components

import "errors"

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
	c := Cell{Mark: NoMark}
	return &c
}

//GetMark returns the Mark of the cell
func (c *Cell) GetMark() string {
	return c.Mark
}

//SetMark returns error if cell is already marked, else it marks the cell
func (c *Cell) SetMark(m string) error {
	if c.Mark == "-" {
		c.Mark = m
		return nil
	}
	return errors.New("\t\t\tCell has been marked already, try an empty cell")
}
