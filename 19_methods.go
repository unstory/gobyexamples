package main

import "fmt"

type rect struct {
	weight, height int
}

func (r rect) area() int { // input a copy
	return r.weight * r.height
}

func (r *rect) perim() int { // input a refference
	return 2 * r.height * r.weight
}

func (r *rect) changeHeightByRef(h int) {
	r.height = h
}

func (r rect) changeHeight(h int) {
	r.height = h
}

func main() {
	rect_1 := rect{weight: 3, height: 4}
	ref_rect := &rect_1 // 使用指针可以避免拷贝，改变接收的值
	fmt.Println("rect_1.area():", rect_1.area())
	fmt.Println("ref_rect.area():", ref_rect.area())

	fmt.Println("rect_1.perim():", rect_1.perim())
	fmt.Println("ref_rect.perim():", ref_rect.perim())

	rect_1.changeHeight(11)
	fmt.Println("rect_1.changeHeight(11):", rect_1.height)
	ref_rect.changeHeightByRef(11)
	fmt.Println("ref_rect.changeHeightByRef(11):", ref_rect.height)

}
