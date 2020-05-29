package components

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		name string
		mark string
	}{
		{"Reigns", OMark},
		{"Lesnar", OMark},
		{"Strowman", XMark},
		{"Joe", XMark},
	}
	for _, i := range tests {
		currentName := i.name
		requiredMark := i.mark
		testPlayer := NewPlayer(currentName, requiredMark)
		returnedName := testPlayer.Name
		returnedMark := testPlayer.Mark
		if currentName != returnedName {
			t.Error(currentName, returnedName)
		}
		if requiredMark != returnedMark {
			t.Error(requiredMark, returnedMark)
		}
	}
}
