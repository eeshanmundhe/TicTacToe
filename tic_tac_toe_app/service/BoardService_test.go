package service

import (
	"errors"
	"testing"
	"tic_tac_toe_app/components"
)

func TestPutMarkInPosition(t *testing.T) {
	tests := []struct {
		wantBoard    *BoardService
		wantPlayer   components.Player
		wantPosition uint8
		wantError    error
	}{
		{&BoardService{components.NewBoard(uint8(3))}, components.Player{Name: "Eeshan", Mark: components.XMark}, 1, nil},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.NoMark},
			},
			Size: 2,
		},
		}, components.Player{Name: "Eeshan", Mark: components.OMark}, 1, errors.New("Cell is already marked")},
	}

	for _, test := range tests {
		got := test.wantBoard.PutMarkInPosition(&test.wantPlayer, test.wantPosition)
		// fmt.Println(test.wantBoard.PrintBoard())
		if test.wantError != nil || got != nil {
			if got.Error() != test.wantError.Error() {
				t.Error(got, test.wantError)
			}
		}
	}
}

func TestPrintBoard(t *testing.T) {
	tests := []struct {
		wantBoard        *BoardService
		wantMatrixString string
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
			Size: 3,
		},
		}, "\n\t\t\t---\n\t\t\tX--\n\t\t\t-X-\n\t\t\t"},
	}
	for _, test := range tests {
		got := test.wantBoard.PrintBoard()
		if test.wantMatrixString != got {
			t.Error(test.wantMatrixString, got)
		}
	}
}

func TestCheckBoardIsFull(t *testing.T) {
	tests := []struct {
		wantBoard *BoardService
		want      bool
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
		got := test.wantBoard.CheckBoardIsFull()
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}
