package service

import (
	"errors"
	"testing"
	"tic_tac_toe_app/components"
)

func TestBeginGame(t *testing.T) {
	var list = []struct {
		position uint8
		g        *GameService
		expected error
	}{
		{10, &GameService{&ResultService{&BoardService{&components.Board{
			Dimension: 3,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			}}}}, &components.Player{Marking: "O", Name: "Roman"}, &components.Player{Marking: "X", Name: "Reigns"}}, errors.New("Positoin is out of bounds")},
	}

	for _, str := range list {
		_, err := str.g.BeginGame(str.position, &components.Player{Name: "AEW", Marking: "X"})
		if err.Error() != str.expected.Error() {
			t.Error(err, str.expected)
		}
	}
}
