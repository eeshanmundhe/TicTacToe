package components

import "testing"

func TestNewBoard(t *testing.T) {
	tests := []struct {
		dim   uint8
		total uint8
		mark  string
	}{
		{4, 16, NoMark},
		{3, 9, NoMark},
		{0, 0, NoMark},
		{5, 25, NoMark},
	}
	for _, i := range tests {
		output := NewBoard(i.dim).Cells
		dimension := uint8(len(output))
		if i.total != dimension {
			t.Error(i.total, dimension)
		}
		for _, j := range output {
			gotMark := j.GetMark()
			if i.mark != gotMark {
				t.Error(i.mark, gotMark)
			}

		}
	}

}
