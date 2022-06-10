package stringreversal

import "strings"

func StringReversal(word string) string {
	splitStr := strings.Split(word, "")
	tmp := []string{}

	for i := len(splitStr) - 1; i >= 0; i-- {
		tmp = append(tmp, splitStr[i])
	}
	return strings.Join(tmp, "")
}
