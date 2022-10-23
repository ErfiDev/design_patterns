package main

import "fmt"

/*
Visitor is a behavioral design pattern that allows
adding new behaviors to existing class hierarchy
without altering any existing code.
*/

type Shape interface {
	accept(v Visitor)
}

type Visitor interface {
	visitForSquare(s *Square)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

type Area struct{}

func (a *Area) visitForSquare(s *Square) {
	fmt.Println("calculating area of square")
}

func main() {
	square := new(Square)
	square.side = 10

	areaCalculator := new(Area)

	square.accept(areaCalculator)

}
