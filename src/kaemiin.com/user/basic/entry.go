package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func swap(x int, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}

func swap2(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println(c1.getArea())
	var a int = 100
	var b int = 200
	fmt.Println(a)
	fmt.Println(b)
	swap(a, b)
	fmt.Println(a)
	fmt.Println(b)
	var c int = 100
	var d int = 200
	fmt.Println(c)
	fmt.Println(d)
	swap2(&c, &d)
	fmt.Println(c)
	fmt.Println(d)
}
