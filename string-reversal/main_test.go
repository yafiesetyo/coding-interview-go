package stringreversal_test

import (
	"testing"

	sr "github.com/yafiesetyo/coding-interview-practise/string-reversal"
)

func TestReverseStringNormalCase(t *testing.T) {
	t.Logf("running ReverseString normal case unit test")

	input := "yafieoawkwk"
	expectedResult := "kwkwaoeifay"

	actualResult := sr.StringReversal(input)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%s' -> Actual '%s'", expectedResult, actualResult)
	}
}

func TestReverseStringWithSpace(t *testing.T) {
	t.Logf("running ReverseString with space unit test")

	input := "yafieoawkwk  "
	expectedResult := "  kwkwaoeifay"

	actualResult := sr.StringReversal(input)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected '%s' -> Actual '%s'", expectedResult, actualResult)
	}
}
