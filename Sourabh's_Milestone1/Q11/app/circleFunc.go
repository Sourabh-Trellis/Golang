package app

type shape interface {
	Area() float32
}

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	area := r.Length * r.Breadth
	return area
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	area := 3.14 * c.Radius * c.Radius
	return area
}
