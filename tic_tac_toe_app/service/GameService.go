package service

import (
	"fmt"
	"tic_tac_toe_app/components"
)

//GameService stores game details
type GameService struct {
	*ResultService
	players [2]*components.Player
}

//NewGameService returns structure GameService
func NewGameService(rs *ResultService, players [2]*components.Player) *GameService {
	return &GameService{rs, players}
}

var turn int

//Play is used to begin game
func (gs *GameService) Play(index uint8) (Result, error) {
	if index < 0 || index >= gs.Size*gs.Size {
		return Result{false, false}, fmt.Errorf("Please enter an integer in range 0-%d", (gs.Size*gs.Size)-1)
	}
	var result Result
	if turn%2 == 0 {
		err := gs.PutMarkInPosition(gs.players[0], index)
		if err != nil {
			return Result{false, false}, err
		}
		result = gs.GetResult(gs.players[0].Mark)
	} else if turn%2 == 1 {
		err := gs.PutMarkInPosition(gs.players[1], index)
		if err != nil {
			return Result{false, false}, err
		}
		result = gs.GetResult(gs.players[1].Mark)
	}
	// fmt.Println("turn", turn)
	turn++
	return result, nil
}
