package fizzbuzz_test

import (
	"strings"
	"testing"

	fb "github.com/yafiesetyo/coding-interview-practise/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	input := 15
	expectedResult := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}

	actualResult := fb.FizzBuzz(input)

	if strings.Join(actualResult, ",") != strings.Join(expectedResult, ",") {
		t.Errorf("Wrong result. Expected '%v' -> Actual '%v'", expectedResult, actualResult)
	}
}
