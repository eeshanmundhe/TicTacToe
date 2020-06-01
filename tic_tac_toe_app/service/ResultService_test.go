package service

import (
	"Tic-Tac-Toe/components"
	"testing"
)

func TestCheckRows(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
			Size: 3,
		},
		},
		}, components.OMark, true},
	}

	for _, test := range tests {
		got := test.rs.checkRow(test.mark)
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}

func TestCheckColumns(t *testing.T) {
	tests := []struct {
		rs   *ResultService
		mark string
		want bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
			Size: 3,
		},
		},
		}, components.OMark, true},
	}

	for _, test := range tests {
		got := test.rs.checkColumn(test.mark)
		if test.want != got {
			t.Error(test.want, got)
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
