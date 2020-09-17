package main

import (
	"errors"
	"fmt"
)

func f1(args int) (int, error) {
	if args == 42 {
		return -1, errors.New("can't work with 42")
	}
	return args + 3, nil
}

//custom define error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// test f1
	for _, i := range []int{7, 42} {
		r, e := f1(i)
		if e != nil {
			fmt.Println("f1 fail, e:", e)
		} else {
			fmt.Println("f1 work, r:", r)
		}
	}
	// test f2
	for _, i := range []int{7, 42} {
		r, e := f2(i)
		if e != nil {
			fmt.Println("f2 fail, e:", e)
		} else {
			fmt.Println("f2 work, r:", r)
		}
	}
	//custom error
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
