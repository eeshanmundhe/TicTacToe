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

//newCell returns the mark from Cell struct
func newCell() *Cell {
	c := Cell{mark: NoMark}
	return &c
}

//getMark returns the mark of the cell
func (c *Cell) getMark() string {
	return c.mark
}

//setMark returns the selected cell with the required mark
func (c *Cell) setMark(m string) error {
	if c.mark == "-" {
		c.mark = m
		return nil
	}
	return errors.New("Cell has been marked")
}
