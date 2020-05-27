package components

import (
	"testing"
)

func TestNewCell(t *testing.T) {
	current := newCell()
	required := NoMark
	if current.mark != required {
		t.Error(current.mark, required)
	}
}

func TestSetMark(t *testing.T) {
	current1 := &Cell{mark: NoMark}
	req1 := "X"
	var err1 error
	err1 = nil
	if current1.setMark(req1) != err1 {
		t.Error(current1.setMark(req1), err1)
	}

	current2 := &Cell{mark: NoMark}
	req2 := "O"
	var err2 error
	err2 = nil
	if current2.setMark(req2) != err2 {
		t.Error(current1.setMark(req2), err2)
	}

	current3 := &Cell{mark: XMark}
	req3 := "X"
	var err3 error
	err3 = nil
	if current3.setMark(req3) != err3 {
		t.Error(current1.setMark(req3), err3)
	}

	current4 := &Cell{mark: XMark}
	req4 := "O"
	var err4 error
	err4 = nil
	if current4.setMark(req4) != err4 {
		t.Error(current4.setMark(req4), err4)
	}

	current5 := &Cell{mark: OMark}
	req5 := "O"
	var err5 error
	err5 = nil
	if current5.setMark(req5) != err5 {
		t.Error(current1.setMark(req5), err5)
	}

	current6 := &Cell{mark: OMark}
	req6 := "X"
	var err6 error
	err6 = nil
	if current6.setMark(req6) != err6 {
		t.Error(current1.setMark(req6), err6)
	}

}

func TestGetMark(t *testing.T) {
	current1 := &Cell{mark: XMark}
	required1 := XMark
	if current1.getMark() != required1 {
		t.Error(current1, required1)
	}

	current2 := &Cell{mark: XMark}
	required2 := XMark
	if current2.getMark() != required2 {
		t.Error(current2, required2)
	}

	current3 := &Cell{mark: XMark}
	required3 := XMark
	if current3.getMark() != required3 {
		t.Error(current3, required3)
	}
}
