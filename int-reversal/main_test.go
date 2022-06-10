package intreversal_test

import (
	"testing"

	ir "github.com/yafiesetyo/coding-interview-practise/int-reversal"
)

func TestIntReversalNormalCase(t *testing.T) {
	input := 521
	expectedResult := 125

	actualResult := ir.IntReversal(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}

func TestIntReversalNegativeNum(t *testing.T) {
	input := -34
	expectedResult := -43

	actualResult := ir.IntReversal(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}

func TestIntReversalWithZero(t *testing.T) {
	input := 1100
	expectedResult := 11

	actualResult := ir.IntReversal(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}

func TestIntReversalWithNegativeAndZero(t *testing.T) {
	input := -910
	expectedResult := -19

	actualResult := ir.IntReversal(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}
