package components

//Board stores Size and cells of the board
type Board struct {
	Cells []*Cell
	Size  uint8
}

//NewBoard returns struct Board
func NewBoard(size uint8) *Board {
	BoardSize := size * size
	var cellArray = make([]*Cell, BoardSize)
	for i := range cellArray {
		cellArray[i] = NewCell()
	}
	return &Board{
		Cells: cellArray,
		Size:  size,
	}
}
