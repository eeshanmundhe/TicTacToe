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
	k := 0
	for i := 0; i < int(r.Dimension*r.Dimension); i++ {
		if r.Cells[i].GetMark() == mark {
			k++
		}
		if k == int(r.Dimension) {
			return true
		}
		if i%int(r.Dimension) == int(r.Dimension)-1 {
			k = 0
		}
	}
	return false
}
func (r *ResultService) checkColumn(mark string) bool {
	//for columns
	kcol := make([]int, r.Dimension)
	for i := 0; i < int(r.Dimension*r.Dimension); i++ {

		if r.Cells[i].GetMark() == mark {
			kcol[i%int(r.Dimension)]++
		}
		if contains(kcol, int(r.Dimension)) {
			return true
		}
	}
	return false
}
func (r *ResultService) checkAcrossLeft(mark string) bool {
	//for LR diagonal
	k := 0
	j := 0
	for i := 0; i < int(r.Dimension); i++ {

		if r.Cells[(i*int(r.Dimension))+j].GetMark() == mark {
			k++
		}
		if k == int(r.Dimension) {
			return true
		}
		j++
	}
	return false
}
func (r *ResultService) checkAcrossRight(mark string) bool {
	//for RL diagonal
	k := 0
	for i := 0; i < int(r.Dimension); i++ {
		if r.Cells[(i+1)*(int(r.Dimension)-1)].GetMark() == mark {
			k++
		}
		if k == int(r.Dimension) {
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
	} else if r.checkAcrossLeft(player.Marking) {
		return "win"
	} else if r.checkAcrossRight(player.Marking) {
		return "win"
	} else if r.CheckGameOver() {
		return "draw"
	}
	return "ongoing"
}
