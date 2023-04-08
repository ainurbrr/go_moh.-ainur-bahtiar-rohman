package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3{
		t.Error("Expected 1 (+) 2 to equal 3")
	}

	if Addition(-1, -2) != -3{
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubstraction(t *testing.T) {
	if Substraction(2, 1) != 1{
		t.Error("Expected 2 (-) 1 to equal 1")
	}

	if Substraction(-1, -2) != 1{
		t.Error("Expected -1 (-) -2 to equal 1")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(2, 2) != 4{
		t.Error("Expected 2 (*) 2 to equal 4")
	}

	if Multiplication(-3, -2) != 6{
		t.Error("Expected -3 (*) -2 to equal 6")
	}
}

func TestDivision(t *testing.T) {
	if Division(2, 2) != 1{
		t.Error("Expected 2 (/) 2 to equal 1")
	}

	if Division(-6, -2) != 3{
		t.Error("Expected -6 (/) -2 to equal 3")
	}
}