package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, i)
	}
}

func main() {
	f("direct")

	go f("go routine")

	go func(tmp string) {
		for j := 0; j < 10; j++ {
			fmt.Println(tmp)
		}
	}("one")
	time.Sleep(time.Second)

	fmt.Println("done")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
