package service

import (
	"testing"
	"tic_tac_toe_app/components"
)

func TestCheckRows(t *testing.T) {
	tests := []struct {
		input    *ResultService
		input1   string
		input2   uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, 4, true},
	}

	for _, test := range tests {
		if test.input.checkRows(test.input1, test.input2) != test.expected {
			t.Error("Check Rows Failed")
		}
	}
}

func TestCheckColumns(t *testing.T) {
	tests := []struct {
		input    *ResultService
		input1   string
		input2   uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, 5, false},
	}

	for _, test := range tests {
		if test.input.checkColumns(test.input1, test.input2) != test.expected {
			t.Error("Check Columns Failed")
		}
	}
}

func TestCheckFirstDiagonal(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
			},
			Size: 3,
		},
		},
		}, components.XMark, true},
	}

	for _, test := range tests {
		got := test.rs.checkFirstDiagonal(test.mark)
		test.rs.checkFirstDiagonal(test.mark)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}

func TestCheckSecondDiagonal(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
			},
			Size: 3,
		},
		},
		}, components.XMark, true},
	}

	for _, test := range tests {
		got := test.rs.checkSecondDiagonal(test.mark)
		test.rs.checkSecondDiagonal(test.mark)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}
