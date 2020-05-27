package components

import (
	"testing"
)

func TestNewCell(t *testing.T) {
	//tests if cell is unmarked initially
	current := newCell()
	required := NoMark
	if current.mark != required {
		t.Error(current.mark, required)
	}
}

func TestSetMark(t *testing.T) {
	//tests the case when cell is not marked and user chooses X
	current1 := &Cell{mark: NoMark}
	req1 := "X"
	var err1 error
	err1 = nil
	if current1.setMark(req1) != err1 {
		t.Error(current1.setMark(req1), err1)
	}

	//tests the case when cell is not marked and user chooses O
	current2 := &Cell{mark: NoMark}
	req2 := "O"
	var err2 error
	err2 = nil
	if current2.setMark(req2) != err2 {
		t.Error(current1.setMark(req2), err2)
	}

	//tests the case when cell is marked X and user chooses X
	current3 := &Cell{mark: XMark}
	req3 := "X"
	var err3 error
	err3 = nil
	if current3.setMark(req3) != err3 {
		t.Error(current1.setMark(req3), err3)
	}

	//tests the case when cell is marked X and user chooses O
	current4 := &Cell{mark: XMark}
	req4 := "O"
	var err4 error
	err4 = nil
	if current4.setMark(req4) != err4 {
		t.Error(current4.setMark(req4), err4)
	}

	//tests the case when cell is marked O and user chooses O
	current5 := &Cell{mark: OMark}
	req5 := "O"
	var err5 error
	err5 = nil
	if current5.setMark(req5) != err5 {
		t.Error(current1.setMark(req5), err5)
	}

	//tests the case when cell is marked O and user chooses X
	current6 := &Cell{mark: OMark}
	req6 := "X"
	var err6 error
	err6 = nil
	if current6.setMark(req6) != err6 {
		t.Error(current1.setMark(req6), err6)
	}

}

func TestGetMark(t *testing.T) {
	//tests if X is marked
	current1 := &Cell{mark: XMark}
	required1 := XMark
	if current1.getMark() != required1 {
		t.Error(current1, required1)
	}

	//tests if O is marked
	current2 := &Cell{mark: OMark}
	required2 := OMark
	if current2.getMark() != required2 {
		t.Error(current2, required2)
	}

	//tests if cell is unmarked
	current3 := &Cell{mark: NoMark}
	required3 := NoMark
	if current3.getMark() != required3 {
		t.Error(current3, required3)
	}
}
