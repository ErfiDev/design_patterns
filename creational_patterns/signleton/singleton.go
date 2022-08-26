package main

import (
	"fmt"
	"sync"
)

/*
Singleton is a creational design pattern that lets
you ensure that a class has only one instance,
while providing a global access point to this instance.
*/

// Singleton pattern solves two problems at the same time.

/*
 1:
	Ensure that a class has just a single instance.
	Why would anyone want to control how many instances
	a class has? The most common reason for this is to
	control access to some shared resource—for example,
	a database or a file.

	Here’s how it works: imagine that you created an
	object, but after a while decided to create a new
	one. Instead of receiving a fresh object, you’ll
	get the one you already created.
*/

/*
 2:
	Provide a global access point to that instance.
	Remember those global variables that you used to
	store some essential objects? While they’re very
	handy, they’re also very unsafe since any code can
	potentially overwrite the contents of those
	variables and crash the app.

	Just like a global variable, the Singleton pattern
	lets you access some object from anywhere in the
	program. However, it also protects that instance
	from being overwritten by other code.
*/

type DB struct{}

var lock = new(sync.Mutex)
var instance *DB

// functionality
func New() *DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating instance...")
			instance = &DB{}
		} else {
			fmt.Println("Instance already initialized!")
		}
	} else {
		return instance
	}

	return instance
}

func main() {
	for i := 0; i <= 10; i++ {
		go New()
	}

	fmt.Scanln()
}
