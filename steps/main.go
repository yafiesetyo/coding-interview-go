package steps

import (
	"fmt"
	"math"
)

func Steps(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func ReverseSteps(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j > n-i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func RecursiveSteps(n int, i int, lvl string) string {
	if i == n {
		return ""
	}

	if n == len(lvl) {
		fmt.Println(lvl)
		return RecursiveSteps(n, i+1, "")
	}

	if len(lvl) <= i {
		lvl += "#"
	} else {
		lvl += " "
	}

	return RecursiveSteps(n, i, lvl)
}

func PyramidSteps(n int) string {
	col := n + (n - 1)
	median := math.Floor(float64(col) / 2)

	for i := 0; i < n; i++ {
		lvl := ""
		for j := 0; j < col; j++ {
			if float64(j) >= median-float64(i) && float64(j) <= median+float64(i) {
				lvl += "*"
			} else {
				lvl += " "
			}
		}
		fmt.Println(lvl)
	}

	return ""
}

func RecursivePyramidSteps(n int, i int, j int, lvl string) string {
	col := n + (n - 1)
	median := math.Floor(float64(col) / 2)

	// main loop
	if i == n {
		return ""
	}

	// inside loop
	if j >= col {
		fmt.Println(lvl)
		return RecursivePyramidSteps(n, i+1, 0, "")
	}

	if float64(j) >= median-float64(i) && float64(j) <= median+float64(i) {
		lvl += "*"
	} else {
		lvl += " "
	}

	return RecursivePyramidSteps(n, i, j+1, lvl)
}
