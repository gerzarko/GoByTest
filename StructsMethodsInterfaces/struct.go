package StructsMethodsInterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {

	return r.Height * r.Width

}

type Circle struct {
	Radius float64
}

type Shape interface {
    Area() float64

}





func (c Circle) Area() float64 {

	return c.Radius * c.Radius * math.Pi

}

func Perimeter(recta Rectangle) float64 {

	totalPerimeter := (recta.Height * 2) + (recta.Width * 2)
	return totalPerimeter

}

func Area(recta Rectangle) float64 {
	return recta.Height * recta.Width
}
