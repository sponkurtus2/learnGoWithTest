package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radious float64
}

type Sphere struct {
	Radious float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radious * c.Radious
}

func (s Sphere) Area() float64 {
	return 4 * math.Pi * math.Pow(s.Radious, 2)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
