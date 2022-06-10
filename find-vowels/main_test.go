package findvowels_test

import (
	"testing"

	fv "github.com/yafiesetyo/coding-interview-practise/find-vowels"
)

func TestFindVowelsNormalCase(t *testing.T) {
	input := "Halo Semuaaaa!!!"
	expectedResult := 8

	actualResult := fv.FindVowels(input)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected %d -> Actual %d", expectedResult, actualResult)
	}

}
