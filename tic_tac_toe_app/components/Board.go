package components

//Board stores cell
type Board struct {
	Matrix    []*Cell
	dimension int
}

//CreateBoard will create a new Board
func CreateBoard(dim int) *Board {
	cellMat := make([]*Cell, dim*dim)
	for i := 0; i < dim*dim; i++ {
		cellMat[i] = NewCell()
	}
	return &Board{
		Matrix:    cellMat,
		dimension: dim,
	}
}
