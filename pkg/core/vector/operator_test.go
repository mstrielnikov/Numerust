package vector

import "testing"

var (
	operatorTestCasesX = [][]int8{
		{1, 1, 0, 1},
		{1, 1, 0, 1, 1, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 0, 1},
	}
	operatorTestCasesY = [][]int8{
		{0, 1, 1, 0},
		{1, 1, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{1, 1, 0, 1, 1, 0},
	}
)

func TestEquals(t *testing.T) {
	expected := []bool{false, true, false, false}
	for i, got := range expected {
		got = Equals(operatorTestCasesX[i], operatorTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestGreater(t *testing.T) {
	expected := []bool{true, false, true, true}
	for i, got := range expected {
		got = Greater(operatorTestCasesX[i], operatorTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestGreaterEquals(t *testing.T) {
	expected := []bool{true, true, true, true}
	for i, got := range expected {
		got = GreaterEquals(operatorTestCasesX[i], operatorTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestLesser(t *testing.T) {
	expected := []bool{false, true, false, true}
	for i, got := range expected {
		got = Lesser(operatorTestCasesX[i], operatorTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestLesserEquals(t *testing.T) {
	expected := []bool{false, true, false, false}
	for i, got := range expected {
		got = LesserEquals(operatorTestCasesX[i], operatorTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], "-- got : ", got)
		}
	}
}