package shapes

import "fmt"

type Reactangle struct { // Reactangle
	width  float64
	height float64
}

func NewReactangle(width float64, height float64) *Reactangle {
	return &Reactangle{width, height}
}

func (r *Reactangle) Area() float64 {
	return r.height * r.width
}

func (r *Reactangle) Display() {
	fmt.Println("height: ", r.height, "\nwidth: ", r.width)
}
