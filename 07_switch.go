package main

import (
	"fmt"
	"time"
)

func main() {
	// 所有分支不匹配后再执行default,有分支匹配了则不运行default
	i := 4
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	wd := time.Now().Weekday()
	switch wd {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is the weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}
}
