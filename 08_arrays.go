package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	var b [10]int
	fmt.Println("b:", b)
	b[4] = 100
	fmt.Println("set b[4]:", b)
	fmt.Println("get b[4]", b[4])
	fmt.Println("len len(b)", len(b))

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("c", c)

	var x [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			x[i][j] = i + j
		}
	}
	fmt.Println("2d:", x)

}
