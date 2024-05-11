package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func totalAreas(shapes ...Shape) float64 {
	area := 0.0
	for _, s := range shapes {
		area += s.area()
	}

	return area
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	c := new(Circle)
	c.area()

	a := new(Android)
	a.Name = "Andy"
	a.talk()

	fmt.Println(totalAreas(c))
}
