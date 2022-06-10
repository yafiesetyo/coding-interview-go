package palindrome_test

import (
	"testing"

	pl "github.com/yafiesetyo/coding-interview-practise/palindrome"
)

func TestPalindromeNormalCase(t *testing.T) {
	input := "abba"
	expectedResult := true

	actualResult := pl.Palindrome(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}

func TestPalindromeWithSpace(t *testing.T) {
	input := "ab  ba"
	expectedResult := true

	actualResult := pl.Palindrome(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}

func TestPalindromeWithWrongInput(t *testing.T) {
	input := "abcdba"
	expectedResult := false

	actualResult := pl.Palindrome(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}

func TestPalindromeWithOneDiffChar(t *testing.T) {
	input := "abcba"
	expectedResult := true

	actualResult := pl.Palindrome(input)

	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}
