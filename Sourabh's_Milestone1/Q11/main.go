package main

import "fmt"

type shape interface {
	area() float32
}

type Rectangle struct {
	length  float32
	breadth float32
}

func (r Rectangle)area() float32 {
	area := r.length * r.breadth
	return area
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	area := 3.14 * c.radius * c.radius
	return area
}

func main() {

	var circle shape = Circle{5}
	fmt.Println("Area of circle is", circle.area())

	var rect shape = Rectangle{10,20}
	fmt.Println("Area of Rectangle is", rect.area())

}
