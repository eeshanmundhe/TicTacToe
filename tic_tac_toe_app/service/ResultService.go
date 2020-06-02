package service

import "tic_tac_toe_app/components"

//ResultService stores BoardService
type ResultService struct {
	*BoardService
}

//Result stores boolean values for win or draw to deermine at the end of game
type Result struct {
	CurrResult string
	Win        bool
	Draw       bool
}

//NewResultService return struct ResultService
func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{boardService}
}

func (res *ResultService) checkRows(mark string, pos uint8) bool {
	ret := true
	i := (pos - pos%res.Size)
	for j := i; j < (i + res.Size); j++ {
		if res.Cells[j].GetMark() != mark {
			ret = false
		}
	}
	return ret
}

func (res *ResultService) checkColumns(mark string, pos uint8) bool {
	ret := true
	for j := pos % res.Size; j < (res.Size * res.Size); j = j + res.Size {
		if res.Cells[j].GetMark() != mark {
			ret = false
		}
	}
	return ret
}

func (res *ResultService) checkFirstDiagonal(mark string) bool {
	size := res.Size
	for i := uint8(0); i < size; i++ {
		if res.Board.Cells[size*i+i].GetMark() != mark {
			return false
		}
	}
	return true
}

func (res *ResultService) checkSecondDiagonal(mark string) bool {
	size := res.Size
	for i := uint8(0); i < size; i++ {
		if res.Board.Cells[(size*i)+(size-1-i)].GetMark() != mark {
			return false
		}
	}
	return true
}

//GetResult return the result of the game
func (res *ResultService) GetResult(player *components.Player, pos uint8) Result {

	if res.checkRows(player.Mark, pos) || res.checkColumns(player.Mark, pos) || res.checkFirstDiagonal(player.Mark) || res.checkFirstDiagonal(player.Mark) {
		return Result{res.PrintResult(player.Name + " Won"), true, false}
	} else if res.CheckBoardIsFull() {
		return Result{res.PrintResult("Draw"), false, true}
	}
	return Result{"", false, false}
}

//PrintResult prints the status of the result
func (res *ResultService) PrintResult(resultStat string) string {
	retString := ""

	for i := 0; i < 3; i++ {
		retString += "\n\t"
		for j := 0; j < len(resultStat)+6; j++ {
			if i == 0 || i == 2 {
				retString += "-"
			}
		}
		if i == 1 {
			retString += "|  " + resultStat + "  |"
		}
	}
	retString += "\n"
	return retString

}
