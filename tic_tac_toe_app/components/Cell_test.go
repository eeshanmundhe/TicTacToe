package components

import (
	"errors"
	"testing"
)

func TestNewCell(t *testing.T) {
	want := NoMark
	got := NewCell()
	if want != got.Mark {
		t.Error(want, got.Mark)
	}
}

func TestMark(t *testing.T) {
	tests := []struct {
		cell *Cell
		mark string
		want error
	}{
		{&Cell{Mark: NoMark}, XMark, nil},
		{&Cell{Mark: NoMark}, OMark, nil},
		{&Cell{Mark: XMark}, OMark, errors.New("Cell is already marked")},
		{&Cell{Mark: XMark}, XMark, errors.New("Cell is already marked")},
		{&Cell{Mark: OMark}, OMark, errors.New("Cell is already marked")},
		{&Cell{Mark: OMark}, XMark, errors.New("Cell is already marked")},
	}
	for _, test := range tests {
		got := test.cell.SetMark(test.mark)
		if test.want != nil || got != nil {
			if test.want.Error() != got.Error() {
				t.Error(test.want, got)
			}
		}
	}
}

func TestGetMark(t *testing.T) {
	tests := []struct {
		cell *Cell
		want string
	}{
		{&Cell{Mark: NoMark}, NoMark},
		{&Cell{Mark: OMark}, OMark},
		{&Cell{Mark: XMark}, XMark},
	}
	for _, test := range tests {
		got := test.cell.GetMark()
		if test.want != got {
			t.Error(test.want, got)
		}
	}

}
