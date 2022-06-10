package intreversal

import (
	"strconv"
	"strings"
)

func IntReversal(input int) int {
	checkPositive := input >= 0
	strInput := strconv.Itoa(input)
	strInputSplit := strings.Split(strInput, "")

	tmp := []string{}

	if !checkPositive {
		tmp = append(tmp, "-")
	}

	for i := len(strInputSplit) - 1; i >= 0; i-- {
		if strInputSplit[i] == "-" || strInputSplit[i] == "0" {
			continue
		}
		tmp = append(tmp, strInputSplit[i])
	}

	strRes := strings.Join(tmp, "")
	result, _ := strconv.Atoi(strRes)

	return result
}
