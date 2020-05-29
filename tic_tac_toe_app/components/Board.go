package components

//Board stores cell
type Board struct {
	Cells     []*Cell
	Dimension uint8
}

//NewBoard ill create a new Board
func NewBoard(dim uint8) *Board {
	TotalCells := dim * dim
	var boardMatrix = make([]*Cell, TotalCells)
	boardMatrix = make([]*Cell, TotalCells)
	for i := range boardMatrix {
		boardMatrix[i] = NewCell()
	}
	return &Board{
		Cells:     boardMatrix,
		Dimension: dim,
	}
}
