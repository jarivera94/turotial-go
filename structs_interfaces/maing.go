package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

//The below func shows how to create methods
func (c *Circle) area() float64 {
	return math.Phi * c.r * c.r
}

type Rectangle struct {
	x, y, x1, y1 float64
}

//Methos to rectangle
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y1)
	w := distance(r.x1, r.y1, r.x1, r.y1)
	return l * w
}

//functions
func circleArea(circle Circle) float64 {
	return math.Phi * circle.r * circle.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - y1
	b := x1 - y2
	return math.Sqrt(a*a + b*b)
}

//embebedd types

type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("Hi my name is ", p.Name)
}

func (p *Person) setName(name string) {
	p.Name = name
}

type Android struct {
	Model  string
	Person Person
}

//interfaces

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	area := 0.0

	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func main() {

	//different ways to allocate
	var circle1 Circle
	circle2 := new(Circle)
	circle3 := Circle{x: 1, y: 2, r: 2.5}
	rectangle := Rectangle{2, 3, 3, 2}
	fmt.Println(circle1)
	fmt.Println(circle2)
	fmt.Println(circle3)
	fmt.Println("The are of circle3 is ", circleArea(circle3), "it was calculated by  function")
	fmt.Println("The are of circle3 is ", circle3.area(), "it was calculated by method in strcut")
	fmt.Println("The are of recangle is ", rectangle.area(), "it was calculated by method in strcut")
	person := new(Person)
	person.setName("Jael Alexander Rivera Oviedo")
	person.talk()

	fmt.Println("Total area ", totalArea(&circle3, &rectangle))
}
