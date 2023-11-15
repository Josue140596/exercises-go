package structs

import "fmt"

// Any type that implements the area() and perimeter() methods
// is said to implement the shape interface
type shape interface {
	area() float64
	perimeter() float64
}

type Triangle struct {
	A, B, C float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

// if you delete this method
// you'll see this error: rectangle does not implement shape (missing method perimeter)
func (r Rectangle) perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (t Triangle) area() float64 {
	return t.A * t.B * t.C
}

func (t Triangle) perimeter() float64 {
	return t.A + t.B + t.C
}

func PrintShape(s shape) {
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}
