package arraychunk

func ArrayChunk(arrInput []int, divider int) [][]int {
	res := [][]int{}
	i := 0

	for i < len(arrInput) {
		if i+divider > len(arrInput) {
			l := len(arrInput)
			res = append(res, arrInput[i:l])
		} else {
			res = append(res, arrInput[i:i+divider])
		}
		i += divider
	}

	return res
}
