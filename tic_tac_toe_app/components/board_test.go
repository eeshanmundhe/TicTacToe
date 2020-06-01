package components

import "testing"

func TestNewBoard(t *testing.T) {
	tests := []struct {
		size       uint8
		matrixSize uint8
		mark       string
	}{
		{0, 0, NoMark},
		{3, 9, NoMark},
	}
	for _, want := range tests {
		got := NewBoard(want.size).Cells
		gotMatrixSize := uint8(len(got))
		if want.matrixSize != gotMatrixSize {
			t.Error(want.matrixSize, gotMatrixSize)
		}
		for _, cell := range got {
			gotMark := cell.GetMark()
			if want.mark != gotMark {
				t.Error(want.mark, gotMark)
			}

		}
	}

}
