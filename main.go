package main

import (
	"fmt"

	fb "github.com/yafiesetyo/coding-interview-practise/fibonacci"
	sm "github.com/yafiesetyo/coding-interview-practise/spiral-matrix"
	st "github.com/yafiesetyo/coding-interview-practise/steps"
)

func main() {
	// array chunk
	// acRes := ac.ArrayChunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 12)
	// fmt.Println(acRes)

	// st.Steps(5)

	// st.RecursiveSteps(5, 0, "")

	menu := "fib"

	if menu == "py" {
		st.RecursivePyramidSteps(10, 0, 0, "")
	} else if menu == "sp" {
		res := sm.SpiralMatrix(10)
		fmt.Println(res)
	} else {
		fbRes := fb.FibRecursive(25)
		fmt.Println(fbRes)
	}
}
