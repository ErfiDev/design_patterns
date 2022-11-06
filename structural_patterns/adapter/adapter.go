package main

import "fmt"

/*
Adapter is a structural design pattern that allows
objects with incompatible interfaces to collaborate.
*/

// Example with lightning and usb port

type MacComputer interface {
	insertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) insertIntoLightningPort() {
	fmt.Println("insert lightning port")
}

type Windows struct{}

func (w *Windows) insertIntoUsbPort() {
	fmt.Println("insert into usb port")
}

type WindowsAdapter struct {
	w *Windows
}

func (w *WindowsAdapter) insertIntoLightningPort() {
	w.w.insertIntoUsbPort()
	fmt.Println("changing usb port to lightning port")
}

func main() {
	fmt.Println("adapter pattern")
}
