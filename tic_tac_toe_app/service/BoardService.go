package service

import "tic_tac_toe_app/components"

//BoardService structure stores board to make changes
type BoardService struct {
	*components.Board
}

//NewBoardService will return the Board after changes
func NewBoardService(board *components.Board) *BoardService {
	return &BoardService{board}
}

//PutMarkInPosition will mark the given position and return error if already marked
func (bs *BoardService) PutMarkInPosition(p *components.Player, position uint8) error {
	err := bs.Board.Cells[position].SetMark(p.Mark)
	return err
}

//PrintBoard will print the updated board
func (bs *BoardService) PrintBoard() string {
	matrixString := "\n\t\t\t"
	boardSize := uint8(len(bs.Board.Cells))
	for i := uint8(0); i < (boardSize); i++ {
		matrixString += bs.Cells[i].Mark
		if i%(bs.Size) == (bs.Size)-1 {
			matrixString += "\n\t\t\t"
		}
		//matrixString += fmt.Sprint(bs.Board.Cells[i].GetMark())
	}
	return matrixString
}

//CheckBoardIsFull will return false if the board is not full
func (bs *BoardService) CheckBoardIsFull() bool {
	boardSize := uint8(len(bs.Board.Cells))
	for i := uint8(0); i < (boardSize); i++ {
		if bs.Board.Cells[i].GetMark() == components.NoMark {
			return false
		}
	}
	return true
}
