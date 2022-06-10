package fibonacci

func Fibonacci(input int) int {
	collection := []int{}
	for i := 0; i <= input; i++ {
		if i < 2 {
			collection = append(collection, i)
		} else {
			collection = append(collection, collection[i-2]+collection[i-1])
		}
	}

	return collection[input]
}

func FibRecursive(input int) int {
	if input < 2 {
		return input
	}

	return FibRecursive(input-2) + FibRecursive(input-1)

}
