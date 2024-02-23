package structs

import (
	"fmt"
)

// Shape is an interface for shapes
type Shape interface {
	Area() float64
}

// Rectangle is a struct for rectangles
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a struct for circles
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.141592653589793
}

// Triangle is a struct for triangles
type Triangle struct {
	Base      float64
	LeftSide  float64
	RightSide float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return t.Base * t.LeftSide * t.RightSide
}

// Perimeter calculates the perimeter of a rectangle
func Perimeter(shape Shape) float64 {
	switch shape := shape.(type) {
	case Rectangle:
		return 2 * (shape.Width + shape.Height)
	case Circle:
		return 2 * shape.Radius * 3.141592653589793
	case Triangle:
		return shape.Base + shape.LeftSide + shape.RightSide
	default:
		return 0
	}
}

func main() {
	rectangle := Rectangle{10.0, 10.0}
	fmt.Println(Perimeter(rectangle))

	circle := Circle{10}
	fmt.Println(Perimeter(circle))
}
