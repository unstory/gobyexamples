package main

import "fmt"

func main() {
	sum := 0
	a := []int{1, 2, 3, 4, 5}
	for i, nums := range a {
		sum += nums
		fmt.Println("current index is:", i)
	}

	kvs := map[string]int{"a": 2, "b": 3, "c": 5}
	for k, v := range kvs {
		fmt.Printf("%s -> %d \n", k, v)
	}

	for _, c := range "go" {
		fmt.Printf("c is: %d \n", c) // 字母会输出为ascii值
	}

}
