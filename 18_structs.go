package main

import "fmt"

type person struct {
	age  int
	name string
}

func CreateNewPerson(name string) *person { //* 在类型前面，代表指针
	p := person{age: 12, name: name}
	return &p //取变量的地址
}

func main() {
	p1 := person{age: 11, name: "bar"}

	fmt.Println("person(11,'bar'):", p1)

	p := CreateNewPerson("xxx")
	fmt.Println("pointer person:", p)
	fmt.Println("person:", *p) // * 在变量前面，代表取地址映射的值
}
