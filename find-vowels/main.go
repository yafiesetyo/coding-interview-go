package findvowels

import "strings"

var vowels = map[string]bool{
	"a": true,
	"i": true,
	"u": true,
	"e": true,
	"o": true,
}

func FindVowels(input string) (res int) {
	split := strings.Split(strings.TrimSpace(strings.ToLower(input)), "")
	for i := 0; i < len(split); i++ {
		if _, ok := vowels[split[i]]; ok {
			res++
		}
	}
	return
}
