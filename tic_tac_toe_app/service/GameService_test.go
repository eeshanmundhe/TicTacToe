package service

import (
	"errors"
	"testing"
	"tic_tac_toe_app/components"
)

func TestPlay(t *testing.T) {
	tests := []struct {
		gs    *GameService
		index uint8
		want  error
	}{
		{&GameService{&ResultService{&BoardService{&components.Board{[]*components.Cell{nil}, 3}}}, [2]*components.Player{nil, nil}}, 10, errors.New("Please enter an integer in range 0-8")},
	}
	for _, test := range tests {
		_, got := test.gs.Play(test.index)
		if test.want.Error() != got.Error() {
			t.Error(got, test.want, got)
		}
	}
}
