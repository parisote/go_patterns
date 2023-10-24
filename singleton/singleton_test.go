package singleton

import "testing"

func TestGetInstanc(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("bad instance")
	}

	expectedCounter := counter1

	currentCont := counter1.AddOne()

	if currentCont != 1 {
		t.Error("Error")
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("")
	}

	currentCont = counter2.AddOne()

	if currentCont != 2 {
		t.Error("")
	}
}
