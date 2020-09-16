package main

import "fmt"

func main() {
	s := make(map[string]int)
	s["a"] = 1
	s["b"] = 2
	s["c"] = 3
	fmt.Println("s:", s)

	fmt.Println("len(s):", len(s))

	val, is_exist := s["b"]
	fmt.Println("val, is_exist:", val, is_exist)

	delete(s, "d") // does not raise exception
	delete(s, "a")

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
