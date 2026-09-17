package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}

func describe(i any) {
	fmt.Println(i)
}

func describeMore(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("it's an int:", v)
	case string:
		fmt.Println("it's a string:", v)
	case Rectangle:
		fmt.Println("it's a rectangle with area:", v.Area())
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	var s Shape = Rectangle{
		Width:  3,
		Height: 4,
	}

	fmt.Println(s.Area())
	printArea(Rectangle{
		Width:  3,
		Height: 4,
	})

	printArea(Circle{
		Radius: 5,
	})

	describe(42)
	describe("Hello")
	describe(Rectangle{
		Width:  1,
		Height: 2,
	})

	describeMore(42)
	describeMore("hello")
	describeMore(Rectangle{Width: 2, Height: 5})
	describeMore(3.14)
}
