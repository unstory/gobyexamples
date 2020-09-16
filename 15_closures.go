package main

import "fmt"

func nextSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := nextSeq()
	fmt.Printf("current i is %d\n", nextInt())
	fmt.Printf("current i is %d\n", nextInt())
	fmt.Printf("current i is %d\n", nextInt())

	new_nextInt := nextSeq()
	fmt.Printf("new current i is %d", new_nextInt())

}
