package main

import "fmt"

/*
Observer is a behavioral design pattern that allows
some objects to notify other objects about changes in
their state.

The Observer pattern provides a way to subscribe and
unsubscribe to and from these events for any object
that implements a subscriber interface.
*/

type Observer interface {
	update(item string)
	getId() string
}

type Subject interface {
	register(o Observer)
	deregister(o Observer)
	notify()
}

type Item struct {
	name      string
	observers []Observer
	inStock   bool
}

func NewItem(name string) Item {
	return Item{
		name: name,
	}
}

func (i *Item) availableIt() {
	i.inStock = true
	fmt.Println("item are availabe in the store!")
	i.notify()
}

func (i *Item) register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) deregister(o Observer) {
	fakeList := i.observers
	obLength := len(fakeList)

	for i, ob := range fakeList {
		if o.getId() == ob.getId() {
			fakeList[obLength-1], fakeList[i] = fakeList[i], fakeList[obLength-1]
		}
	}

	i.observers = fakeList[:obLength-1]
}

func (i *Item) notify() {
	for _, ob := range i.observers {
		ob.update(i.name)
	}
}

type Customer struct {
	id string
}

func (c *Customer) update(item string) {
	fmt.Printf("sending %s to customer id %s\n", item, c.id)
}

func (c *Customer) getId() string {
	return c.id
}

func main() {
	// new item
	carpet := NewItem("carpet")

	// observers
	john := Customer{
		id: "john@mail.com",
	}
	ali := Customer{
		id: "ali@mail.com",
	}

	carpet.register(&john)
	carpet.register(&ali)

	carpet.availableIt()
}
