package service

import (
	"testing"
	"tic_tac_toe_app/components"
)

func TestCheckRow(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Dimension: 2,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, true},

		{&ResultService{&BoardService{&components.Board{
			Dimension: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.OMark, true},
	}
	for _, test := range tests {
		if test.input.checkRow(test.mark) != test.expected {
			t.Error("check row failed")
		}
	}
}

func TestCheckColumn(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Dimension: 2,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.XMark, true},

		{&ResultService{&BoardService{&components.Board{
			Dimension: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.OMark, true},
	}
	for _, test := range tests {
		if test.input.checkColumn(test.mark) != test.expected {
			t.Error("check column failed")
		}
	}
}

func TestCheckLRDiagonal(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Dimension: 2,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, true},

		{&ResultService{&BoardService{&components.Board{
			Dimension: 3,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, true},
	}
	for _, test := range tests {
		if test.input.checkLRDiagonal(test.mark) != test.expected {
			t.Error("check LR diagonal failed")
		}
	}
}

func TestCheckRLDiagonal(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Dimension: 3,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, true},
	}
	for _, test := range tests {
		if test.input.checkRLDiagonal(test.mark) != test.expected {
			t.Error("check RL diagonal failed")
		}
	}
}
