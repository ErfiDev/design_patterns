package main

import "fmt"

/*
Command is a behavioral design pattern that turns a
request into a stand-alone object that contains all
information about the request. This transformation
lets you pass requests as a method arguments, delay or
queue a request’s execution, and support undoable
operations.
*/

type Calculator struct {
	current    float64
	tapHistory []Button
}

func (c *Calculator) tap(val float64, btn Button) {
	res := btn.execute(c.current, val)
	c.current = res
	c.tapHistory = append(c.tapHistory, btn)
}

func (c *Calculator) undo() {
	c.current = c.tapHistory[len(c.tapHistory)-1].undo()
}

type Button interface {
	execute(value, by float64) float64
	undo() float64
}

type AddBtn struct {
	current float64
	history []float64
}

func (a *AddBtn) execute(value, by float64) float64 {
	a.current = value + by
	return a.current
}

func (a *AddBtn) undo() float64 {
	a.current = a.current - a.history[len(a.history)-1 : 0][0]
	a.history = a.history[:len(a.history)-2]
	return a.current
}

func main() {
	calc := new(Calculator)
	addBtn := new(AddBtn)

	calc.tap(10, addBtn)
	calc.tap(15, addBtn)
	fmt.Println(calc.current)
	calc.undo()
	fmt.Println(calc.current)
}
