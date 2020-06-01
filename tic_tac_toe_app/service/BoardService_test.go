package service

import (
	"errors"
	"testing"
	"tic_tac_toe_app/components"
)

func TestPutMarkInPosition(t *testing.T) {
	tests := []struct {
		requiredBoard    *BoardService
		requiredPlayer   components.Player
		requiredPosition uint8
		requiredError    error
	}{
		{&BoardService{components.NewBoard(uint8(3))}, components.Player{Name: "Roman", Marking: components.XMark}, 1, nil},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.NoMark},
			},
			Dimension: 3,
		},
		}, components.Player{Name: "Roman", Marking: components.OMark}, 1, errors.New("Cell has been marked already")},
	}

	for _, test := range tests {
		got := test.requiredBoard.PutMarkInPosition(test.requiredPosition, &test.requiredPlayer)
		if test.requiredError != nil || got != nil {
			if got.Error() != test.requiredError.Error() {
				t.Error(got, test.requiredError)
			}
		}
	}
}

func TestDisplay(t *testing.T) {
	tests := []struct {
		requiredBoard  *BoardService
		requiredMatrix string
	}{
		{&BoardService{components.NewBoard(uint8(4))}, "\n\t\t\t----\n\t\t\t----\n\t\t\t----\n\t\t\t----\n\t\t\t"},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.NoMark},
			},
			Dimension: 3,
		},
		}, "\n\t\t\t---\n\t\t\tX--\n\t\t\t-X-\n\t\t\t"},
	}
	for _, test := range tests {
		got := test.requiredBoard.Display()
		if test.requiredMatrix != got {
			t.Error(test.requiredMatrix, got)
		}
	}
}

func TestCheckGameOver(t *testing.T) {
	tests := []struct {
		requiredBoard *BoardService
		want          bool
	}{
		{&BoardService{components.NewBoard(uint8(3))}, false},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
			},
		},
		}, false},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
			},
		},
		}, true},
	}
	for _, test := range tests {
		got := test.requiredBoard.CheckGameOver()
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}
