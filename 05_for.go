package main

import "fmt"

func main() {
	// 1. for 判断条件
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 2. for初始值，条件，变化
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// 3. 类似于while， break\
	for {
		fmt.Println("loop")
		break
	}
}
