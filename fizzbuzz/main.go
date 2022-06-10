package fizzbuzz

import "strconv"

func FizzBuzz(input int) []string {
	result := []string{}

	for i := 1; i <= input; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "fizzbuzz")
		} else if i%3 == 0 {
			result = append(result, "fizz")
		} else if i%5 == 0 {
			result = append(result, "buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
