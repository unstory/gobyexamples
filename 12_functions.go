package main

import "fmt"

func plus(a int, b int) int {
	c := a + b
	return c

}

func main() {
	c := plus(3, 5)
	fmt.Println(c)
}
