package structs

import (
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	center, radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Sqrt(c.radius);
}

func NewCircle(center, radius float64) Circle {
		return Circle{center, radius};
}

func GetArea(s Shape) float64 {
	return s.area();
}


