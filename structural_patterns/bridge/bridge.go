package main

import (
	"fmt"
	"log"
)

/*
	Bridge pattern is a structural design
	pattern that divides business logic
	or huge class into seprate class
	hierarchies that can be developed
	independently.

	One of these hierarchies (often called the Abstraction)
	will get a reference to an object of the second hierarchy
	some (sometimes, most) of its calls to the implementations object.
	Since all implementations will have a common interface, they’d be
	interchangeable inside the abstraction

	Say, you have two types of computers: Mac and Windows. Also,
	two types of printers: Epson and HP. Both computers and printers
	need to work with each other in any combination. The client doesn’t
	want to worry about the details of connecting printers to computers.

	If we introduce new printers, we don’t want our code to grow exponentially.
 	Instead of creating four structs for the 2*2 combination, we create two
	hierarchies:
	Abstraction hierarchy: this will be our computers
	Implementation hierarchy: this will be our printers
*/

type Computer interface {
	SetPrinter(Printer)
	Print(data string)
}

type Printer interface {
	Print(data string)
}

type Mac struct {
	p Printer
}

func (m *Mac) SetPrinter(p Printer) {
	if p != nil {
		m.p = p
	}

	fmt.Println("printer set!")
}

func (m *Mac) Print(data string) {
	m.p.Print(data)
}

type Windows struct {
	p Printer
}

func (m *Windows) SetPrinter(p Printer) {
	if p != nil {
		m.p = p
	}

	fmt.Println("printer set!")
}

func (m *Windows) Print(data string) {
	m.p.Print(data)
}

type HPPrinter struct{}

func (h *HPPrinter) Print(data string) {
	if data != "" {
		log.Println(data)
	}
}

type CanonPrinter struct{}

func (h *CanonPrinter) Print(data string) {
	if data != "" {
		log.Println(data)
	}
}

func main() {
	win := Windows{}
	mac := Mac{}

	canon := CanonPrinter{}
	HP := HPPrinter{}

	win.SetPrinter(&canon)
	win.Print("print from windows with canon printer")

	win.SetPrinter(&HP)
	win.Print("print from windows with HP printer")

	mac.SetPrinter(&canon)
	mac.Print("print from mac with canon printer")

	mac.SetPrinter(&HP)
	mac.Print("print from mac with HP printer")
}
