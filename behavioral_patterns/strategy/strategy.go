package main

import (
	"errors"
	"fmt"
)

/*
Strategy is a behavioral design pattern that turns a
set of behaviors into objects and makes them
interchangeable inside original context object.
*/

type EvictionContract interface {
	evict(c *CacheStorage)
}

type CacheStorage struct {
	storage        map[string]string
	maxCapacity    int
	capacity       int
	evictionMethod EvictionContract
}

func (c *CacheStorage) set(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *CacheStorage) get(key string) (string, error) {
	value, ok := c.storage[key]
	if !ok {
		return "", errors.New("not found!")
	}

	return value, nil
}

func (c *CacheStorage) evict() {
	c.evictionMethod.evict(c)
	c.capacity--
}

func (c *CacheStorage) changeEvictAlgorithm(e EvictionContract) {
	c.evictionMethod = e
}

// First in, first out evicting method
type Fifo struct{}

func (f *Fifo) evict(c *CacheStorage) {
	fmt.Println("Evicting by fifo method")
}

// Least frequently used
type Lfu struct{}

func (l *Lfu) evict(c *CacheStorage) {
	fmt.Println("Evicting by lfu method")
}

func main() {
	// initializing eviction methods
	lfu := new(Lfu)
	fifo := new(Fifo)

	// cache storage setup
	cs := CacheStorage{
		storage:        make(map[string]string),
		maxCapacity:    100,
		capacity:       0,
		evictionMethod: lfu,
	}

	cs.set("a", "b")
	cs.set("c", "d")
	cs.evict()

	cs.changeEvictAlgorithm(fifo)
	cs.set("e", "f")
	cs.evict()
}
