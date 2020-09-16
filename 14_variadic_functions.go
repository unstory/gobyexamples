package main

import "fmt"

func plus(nums ...int) int {
	sums := 0
	fmt.Println("nums:", nums)
	for _, num := range nums {
		sums += num
	}
	return sums
}

func main() {
	sums := plus(1, 2, 3, 4, 5)
	fmt.Printf("sums: %d", sums)
}
