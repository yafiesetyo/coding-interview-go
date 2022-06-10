package spiralmatrix

func SpiralMatrix(input int) (output [][]int) {

	//generate initial array with 0
	for i := 0; i < input; i++ {
		tmp := []int{}
		for i := 0; i < input; i++ {
			tmp = append(tmp, 0)
		}
		output = append(output, tmp)
	}

	//row pointer
	stR := 0
	edR := input - 1

	//column pointer
	stC := 0
	edC := input - 1

	counter := 1

	for stR <= edR && stC <= edC {
		//left corner top right to left
		for i := stC; i <= edC; i++ {
			output[stR][i] = counter
			counter++
		}
		stR++

		//from end of column (actually on right), go to bottom
		for i := stR; i <= edR; i++ {
			output[i][edC] = counter
			counter++
		}
		edC--

		//from bottom right corner, go to left
		for i := edC; i >= stC; i-- {
			output[edR][i] = counter
			counter++
		}
		edR--

		//from bottom left corner, go up
		for i := edR; i >= stR; i-- {
			output[i][stC] = counter
			counter++
		}
		stC++
	}

	return
}
