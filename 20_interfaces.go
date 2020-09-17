package main

import (
	"fmt"
	"math"
)

type geom interface {
	area() float64
	perim() float64
}

type rect struct {
	width, heigh float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.heigh * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r rect) perim() float64 {
	return r.heigh * r.width * 2.0
}

func (c circle) perim() float64 {
	return 2.0 * math.Pi * c.r
}

func message(g geom) {
	fmt.Println("g:", g)
	a := g.area()
	p := g.perim()
	fmt.Println("area is:", a)
	fmt.Println("perim is:", p)
}

func main() {
	r := rect{heigh: 3, width: 4}
	c := circle{r: 5}
	message(r)
	message(c)
}
