package components

//Board stores cell
type Board struct {
	Matrix    []*Cell
	dimension int
}

//createBoard will create a new Board
func createBoard(dim int) *Board {
	cellMat := make([]*Cell, dim*dim)
	for i := 0; i < dim*dim; i++ {
		cellMat[i] = newCell()
	}
	return &Board{
		Matrix:    cellMat,
		dimension: dim,
	}
}
