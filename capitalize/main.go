package capitalize

import "strings"

func Capitalize(input string) string {
	split := strings.Split(input, " ")
	arrRes := []string{}

	for i := 0; i < len(split); i++ {
		arrRes = append(arrRes, capitalHelper(split[i]))
	}

	return strings.Join(arrRes, " ")
}

func capitalHelper(str string) string {
	split := strings.Split(str, "")
	firstLetter := strings.ToUpper(split[0])
	split[0] = firstLetter

	return strings.Join(split, "")
}
