package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeropointer(i *int) { // *  -->pointer type
	*i = 0 //point to i's value
}

func main() {
	i := 1
	fmt.Println("init i is:", i)
	zeroval(i)
	fmt.Println("zeroval(i):", i)

	j := 1
	fmt.Println("init j is:", j)
	zeropointer(&j) // & --> get address, j's address
	fmt.Println("zeropointer(j):", j)
}
