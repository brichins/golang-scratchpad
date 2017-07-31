package pack

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Log("Failed to add one and two")
		t.Fail()
	}
}

func TestAddVarArgs(t *testing.T) {
	result := Add(1, 2, 3, 4, 5)
	if result != 15 {
		t.Errorf("Failed to add 1-5.  Expected 15: Returned %d", result)
	}
}
