package capitalize_test

import (
	"testing"

	cp "github.com/yafiesetyo/coding-interview-practise/capitalize"
)

func TestCapitalize(t *testing.T) {
	input := "some sentences belongs to what?"
	expectedResult := "Some Sentences Belongs To What?"

	actualResult := cp.Capitalize(input)
	if actualResult != expectedResult {
		t.Errorf("Wrong result. Expected %v -> Actual %v", expectedResult, actualResult)
	}
}
