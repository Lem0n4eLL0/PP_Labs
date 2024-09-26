package shapes

import "fmt"

const PI float64 = 3.1415926535

type Circle struct { // Circle
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

func (c *Circle) Area() float64 {
	return PI * c.radius * c.radius
}

func (c *Circle) Display() {
	fmt.Println("radius: ", c.radius)
}
