package service

//ResultService stores BoardService
type ResultService struct {
	*BoardService
}

//Result stores boolean values for win or draw to deermine at the end of game
type Result struct {
	Win  bool
	Draw bool
}

//NewResultService return struct ResultService
func NewResultService(bs *BoardService) *ResultService {
	return &ResultService{bs}
}

func (rs *ResultService) checkRow(mark string) bool {
	size := rs.Size
	count := 0
	for i := uint8(0); i < size*size; i++ {
		if rs.Board.Cells[i].GetMark() == mark {
			count++
		}
		if count == int(size) {
			return true
		}
		if (i+1)%(size) == 0 {
			count = 0
		}
	}
	return false
}

func (rs *ResultService) checkColumn(mark string) bool {
	size := rs.Size
	for i := uint8(0); i < size; i++ {
		count := 0
		for j := i; j <= (size*size)-(size-i); j += size {
			if rs.Board.Cells[j].GetMark() == mark {
				count++
			}
			if count == int(size) {
				return true
			}
		}
	}
	return false
}

func (rs *ResultService) checkFirstDiagonal(mark string) bool {
	size := rs.Size
	for i := uint8(0); i < size; i++ {
		if rs.Board.Cells[size*i+i].GetMark() != mark {
			return false
		}
	}
	return true
}

func (rs *ResultService) checkSecondDiagonal(mark string) bool {
	size := rs.Size
	for i := uint8(0); i < size; i++ {
		if rs.Board.Cells[(size*i)+(size-1-i)].GetMark() != mark {
			return false
		}
	}
	return true
}

//GetResult return the result of the game
func (rs *ResultService) GetResult(mark string) Result {
	if rs.checkRow(mark) {
		return Result{true, false}
	} else if rs.checkColumn(mark) {
		return Result{true, false}
	} else if rs.checkFirstDiagonal(mark) {
		return Result{true, false}
	} else if rs.checkSecondDiagonal(mark) {
		return Result{true, false}
	} else if rs.CheckBoardIsFull() {
		return Result{false, true}
	}
	return Result{false, false}
}
