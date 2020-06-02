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
func (b *BoardService) PutMarkInPosition(player *components.Player, position uint8) error {
	err := b.Cells[position].SetMark(player.Mark)
	return err
}

//PrintBoard will print the updated board
func (b *BoardService) PrintBoard() string {
	matrixString := "\n\t\t\t"
	boardSize := uint8(len(b.Board.Cells))
	for i := uint8(0); i < (boardSize); i++ {
		matrixString += b.Cells[i].Mark
		if i%(b.Size) == (b.Size)-1 {
			matrixString += "\n\t\t\t"
		}
		//matrixString += fmt.Sprint(b.Board.Cells[i].GetMark())
	}
	return matrixString
}

//CheckBoardIsFull will return false if the board is not full
func (b *BoardService) CheckBoardIsFull() bool {
	for _, cell := range b.Cells {
		if cell.GetMark() == components.NoMark {
			return false
		}
	}
	return true
}
