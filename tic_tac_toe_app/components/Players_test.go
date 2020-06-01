package components

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		name string
		mark string
	}{
		{"Strowman", XMark},
		{"Roman", OMark},
	}
	for _, test := range tests {
		wantName := test.name
		wantMark := test.mark
		testPlayer := NewPlayer(wantName, wantMark)
		gotName := testPlayer.Name
		gotMark := testPlayer.Mark
		if wantName != gotName {
			t.Error(wantName, gotName)
		}
		if wantMark != gotMark {
			t.Error(wantMark, gotMark)
		}
	}
}
