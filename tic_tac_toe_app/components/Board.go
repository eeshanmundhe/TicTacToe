package components

//Board stores cell
type Board struct {
	Cells     []*Cell
	Dimension uint8
}

//NewBoard ill create a new Board
func NewBoard(dim uint8) *Board {
	TotalCells := dim * dim
	var BoardMatrix = make([]*Cell, TotalCells)
	for i := range BoardMatrix {
		BoardMatrix[i] = NewCell()
	}
	return &Board{
		Cells:     BoardMatrix,
		Dimension: dim,
	}
}
