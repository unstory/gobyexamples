package main

import "fmt"

func main(){
	i := 7
	if i % 2 != 0 {
		fmt.Println("i is odd")
	} else {
		fmt.Println("i is even")
	}

	num := 0
	if num > 0 {
		fmt.Println("num is greater than zero")
	} else if num == 0{
		fmt.Println("num is equal to zero")
	} else {
		fmt.Println("num is less than zero")
	}
}