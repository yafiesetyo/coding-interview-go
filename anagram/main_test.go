package anagram_test

import (
	"testing"

	an "github.com/yafiesetyo/coding-interview-practise/anagram"
)

func TestAnagramNormalCase(t *testing.T) {
	input1 := "rail safety"
	input2 := "fairy tales"
	expectedResult := true

	actualResult := an.Anagram(input1, input2)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected %v -> Actual %v", expectedResult, actualResult)
	}
}

func TestAnagramWithNonWordCase(t *testing.T) {
	input1 := "rail! safety!"
	input2 := "fairy tales"
	expectedResult := true

	actualResult := an.Anagram(input1, input2)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected %v -> Actual %v", expectedResult, actualResult)
	}
}

func TestAnagramWithNonWordWithUpperCaseCase(t *testing.T) {
	input1 := "RAIL! SAFETY!"
	input2 := "fairy tales"
	expectedResult := true

	actualResult := an.Anagram(input1, input2)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected %v -> Actual %v", expectedResult, actualResult)
	}
}
func TestAnagramWithNonWordWithUpperCaseBadCase(t *testing.T) {
	input1 := "RAIL! SAFET!"
	input2 := "fairy tales"
	expectedResult := false

	actualResult := an.Anagram(input1, input2)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected %v -> Actual %v", expectedResult, actualResult)
	}
}

func TestAnagramBadCase(t *testing.T) {
	input1 := "RAIL! SAFETY!"
	input2 := "Savety liar"
	expectedResult := false

	actualResult := an.Anagram(input1, input2)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected %v -> Actual %v", expectedResult, actualResult)
	}
}
