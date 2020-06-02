package service

import (
	"fmt"
	"tic_tac_toe_app/components"
)

//GameService stores game details
type GameService struct {
	*ResultService
	players [2]*components.Player
	turn    int
}

//NewGameService returns structure GameService
func NewGameService(res *ResultService, players [2]*components.Player) *GameService {
	return &GameService{res, players, 0}
}

//Play is used to begin game
func (game *GameService) Play(pos uint8) (Result, error) {
	if pos < 0 || pos >= game.Size*game.Size {
		return Result{"", false, false}, fmt.Errorf("Integer should be in range [0,%d]", game.Size*game.Size-1)
	}
	var res Result
	if game.turn%2 == 0 {
		err := game.PutMarkInPosition(game.players[0], pos)
		if err != nil {
			return Result{"", false, false}, err
		}
		res = game.GetResult(game.players[0], pos)
	} else if game.turn%2 == 1 {
		err := game.PutMarkInPosition(game.players[1], pos)
		if err != nil {
			return Result{"", false, false}, err
		}
		res = game.GetResult(game.players[1], pos)
	}
	game.turn++
	return res, nil
}
