package maxchars_test

import (
	"testing"

	mc "github.com/yafiesetyo/coding-interview-practise/max-chars"
)

func TestMaxCharsNormalCase(t *testing.T) {
	input := "aaaaaaaaaaabhhhhhh"
	expectedResult := "a"

	actualResult := mc.MaxChars(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}

func TestMaxCharsWithCapitalizeCase(t *testing.T) {
	input := "Waluyo wiwohowwwwwww"
	expectedResult := "w"

	actualResult := mc.MaxChars(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}
