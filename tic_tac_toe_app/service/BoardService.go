package service

import (
	"tic_tac_toe_app/components"
)

//BoardService structure stores board to make changes
type BoardService struct {
	*components.Board
}

//NewBoardService will return the Board after changes
func NewBoardService(size int) *BoardService {
	return &BoardService{components.NewBoard(uint8(size))}
}

//PutMarkInPosition will mark the given position and return error if already marked
func (b *BoardService) PutMarkInPosition(position uint8, player *components.Player) error {
	err := b.Cells[position].SetMark(player.Marking)
	if err != nil {
		return err
	}
	return nil
}

//PrintBoard will print the updated board
func (b *BoardService) PrintBoard() string {
	out := "\n\t\t\t"
	for i := 0; i < int(b.Dimension*b.Dimension); i++ {
		out += b.Cells[i].Mark
		if i%int(b.Dimension) == int(b.Dimension)-1 {
			out += "\n\t\t\t"
		}
	}
	return out
}

//CheckBoardIsFull will return false if the board is not full
func (b *BoardService) CheckBoardIsFull() bool {
	for i := 0; i < int(b.Dimension*b.Dimension); i++ {
		if b.Cells[i].GetMark() == components.NoMark {
			return false
		}
	}
	return true
}
