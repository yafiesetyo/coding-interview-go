package palindrome

import (
	"strings"
)

func Palindrome(input string) bool {
	tmp := map[string]int{}
	split := strings.Split(input, "")

	for i := 0; i < len(split); i++ {
		if _, ok := tmp[split[i]]; !ok {
			tmp[split[i]] = 1
		} else {
			tmp[split[i]]++
		}
	}

	oddCounter := 0
	for _, v := range tmp {
		if v%2 != 0 {
			oddCounter++
		}
	}

	return oddCounter <= 1
}
