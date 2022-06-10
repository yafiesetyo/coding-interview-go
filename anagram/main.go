package anagram

import (
	"regexp"
	"strings"
)

func Anagram(strA string, strB string) bool {
	// regex n replace
	rgx, _ := regexp.Compile(`\W`)
	strA = rgx.ReplaceAllString(strA, "")
	strB = rgx.ReplaceAllString(strB, "")

	// to lower
	strA = strings.ToLower(strA)
	strB = strings.ToLower(strB)

	// split
	splitA := strings.Split(strA, "")
	splitB := strings.Split(strB, "")

	if len(splitA) != len(splitB) {
		return false
	}

	//def map
	mapA := map[string]int{}
	mapB := map[string]int{}

	for i := 0; i < len(splitA); i++ {
		if _, ok := mapA[splitA[i]]; !ok {
			mapA[splitA[i]] = 1
		} else {
			mapA[splitA[i]]++
		}
	}

	for i := 0; i < len(splitB); i++ {
		if _, ok := mapB[splitB[i]]; !ok {
			mapB[splitB[i]] = 1
		} else {
			mapB[splitB[i]]++
		}
	}

	for k, v := range mapA {
		if _, ok := mapB[k]; !ok {
			return false
		}

		if mapB[k] != v {
			return false
		}
	}

	return true
}
