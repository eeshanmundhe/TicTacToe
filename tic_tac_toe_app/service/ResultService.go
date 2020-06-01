package service

import "tic_tac_toe_app/components"

//ResultService structure helps togenerate result
type ResultService struct {
	*BoardService
}

//NewResultService will return the structure after servicing
func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{boardService}
}

func (r *ResultService) checkRow(mark string) bool {
	//for rows
	count := 0
	for i := 0; i < int(r.Dimension*r.Dimension); i++ {
		if r.Cells[i].GetMark() == mark {
			count++
		}
		if count == int(r.Dimension) {
			return true
		}
		if i%int(r.Dimension) == int(r.Dimension)-1 {
			count = 0
		}
	}
	return false
}
func (r *ResultService) checkColumn(mark string) bool {
	//for columns
	countcol := make([]int, r.Dimension)
	for i := 0; i < int(r.Dimension*r.Dimension); i++ {

		if r.Cells[i].GetMark() == mark {
			countcol[i%int(r.Dimension)]++
		}
		if contains(countcol, int(r.Dimension)) {
			return true
		}
	}
	return false
}
func (r *ResultService) checkLRDiagonal(mark string) bool {
	//for LR diagonal
	count := 0
	j := 0
	for i := 0; i < int(r.Dimension); i++ {

		if r.Cells[(i*int(r.Dimension))+j].GetMark() == mark {
			count++
		}
		if count == int(r.Dimension) {
			return true
		}
		j++
	}
	return false
}
func (r *ResultService) checkRLDiagonal(mark string) bool {
	//for RL diagonal
	count := 0
	for i := 0; i < int(r.Dimension); i++ {
		if r.Cells[(i+1)*(int(r.Dimension)-1)].GetMark() == mark {
			count++
		}
		if count == int(r.Dimension) {
			return true
		}
	}
	return false
}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//GiveResult will give the status of game
func (r *ResultService) GiveResult(player *components.Player) string {

	if r.checkRow(player.Marking) {
		return "win"
	} else if r.checkColumn(player.Marking) {
		return "win"
	} else if r.checkLRDiagonal(player.Marking) {
		return "win"
	} else if r.checkRLDiagonal(player.Marking) {
		return "win"
	} else if r.CheckBoardIsFull() {
		return "draw"
	}
	return "ongoing"
}
