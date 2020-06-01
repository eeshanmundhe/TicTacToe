package service

import (
	"errors"
	"tic_tac_toe_app/components"
)

//GameService stores game details
type GameService struct {
	*ResultService
	player1 *components.Player
	player2 *components.Player
}

//NewGameService returns structure GameService
func NewGameService(pl1, pl2 *components.Player, size int) *GameService {
	return &GameService{NewResultService(NewBoardService(size)), pl1, pl2}
}

//Play is used to play game
func (g *GameService) Play(pos uint8, pl *components.Player) (string, error) {
	if pos < 0 || pos > g.Dimension*g.Dimension-1 {
		return "nil", errors.New("\t\t\tPosition is out of bounds")
	}
	err := g.PutMarkInPosition(pos, pl)
	if err != nil {
		return "nil", err
	}
	return g.GiveResult(pl), nil
}
