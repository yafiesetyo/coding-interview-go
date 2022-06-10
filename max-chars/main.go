package maxchars

import "strings"

func MaxChars(input string) string {
	split := strings.Split(input, "")
	tmp := map[string]int{}

	// populate map
	for i := 0; i < len(split); i++ {
		if _, ok := tmp[split[i]]; !ok {
			tmp[split[i]] = 1
		} else {
			tmp[split[i]]++
		}
	}

	res := ""
	comparator := 0
	for k, v := range tmp {
		if v > comparator {
			comparator = v
			res = k
		}
	}

	return res
}
