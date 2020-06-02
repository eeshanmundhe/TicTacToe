package service

import (
	"errors"
	"testing"
	"tic_tac_toe_app/components"
)

func TestPlay(t *testing.T) {
	tests := []struct {
		input         *GameService
		input1        uint8
		expected1     error
		expectedmark1 string
		input2        uint8
		expected2     error
		expectedmark2 string
	}{
		{&GameService{&ResultService{&BoardService{&components.Board{
			Size: 4,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		}}}, [2]*components.Player{
			{Name: "Brock", Mark: components.OMark},
			{Name: "Ash", Mark: components.XMark},
		}, 1}, 18, errors.New("Integer should be in range [0,15]"), "", 15, nil, components.XMark},

		{&GameService{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		}}}, [2]*components.Player{
			{Name: "Eeshan", Mark: components.OMark},
			{Name: "Sairam", Mark: components.XMark},
		}, 0}, 7, nil, components.OMark, 4, errors.New("Cell is already marked"), ""},
	}
	for _, test := range tests {
		_, err := test.input.Play(test.input1)
		if err != nil {
			if err.Error() != test.expected1.Error() {
				t.Error("Errors don't match", err.Error(), test.expected1.Error())
			}
		} else if test.input.Cells[test.input1].GetMark() != test.expectedmark1 {
			t.Error("Mark didn't match")
		}
		_, err = test.input.Play(test.input2)
		if err != nil {
			if err.Error() != test.expected2.Error() {
				t.Error("Errors don't match", err.Error(), test.expected2.Error())
			}
		} else if test.input.Cells[test.input2].GetMark() != test.expectedmark2 {
			t.Error("Mark didn't match")
		}

	}
}
