package components

import "errors"

//Cell takes mark of string type
type Cell struct {
	mark string
}

const (
	//NoMark stores the mark for unused cell
	NoMark = "-"

	//XMark is used to assign X to the selected Cell
	XMark = "X"

	//OMark is used to assign O to the selected Cell
	OMark = "O"
)

//NewCell returns the mark from Cell struct
func NewCell() *Cell {
	c := Cell{mark: NoMark}
	return &c
}

//GetMark returns the mark of the cell
func (c *Cell) GetMark() string {
	return c.mark
}

//SetMark returns error if cell is already marked, else it marks the cell
func (c *Cell) SetMark(m string) error {
	if c.mark == "-" {
		c.mark = m
		return nil
	}
	return errors.New("Cell has been marked")
}
