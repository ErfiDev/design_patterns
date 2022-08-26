package main

import (
	"errors"
	"fmt"
)

/*
Factory Method is a creational design pattern
that provides an interface for creating objects
in a superclass, but allows subclasses to alter
the type of objects that will be created.
*/

type GunContract interface {
	shoot()
	getType() string
}

type Gun struct {
	Type string
}

func (g *Gun) shoot() {}
func (g *Gun) getType() string {
	return g.Type
}

type Pistol struct {
	Gun
}

type Shotgun struct {
	Gun
}

func NewPistol() *Pistol {
	return &Pistol{
		Gun: Gun{
			Type: "Pistol",
		},
	}
}

func NewShotgun() *Shotgun {
	return &Shotgun{
		Gun: Gun{
			Type: "Shotgun",
		},
	}
}

// Factory method
func New(t string) (GunContract, error) {
	switch t {
	case "Pistol":
		return NewPistol(), nil

	case "Shotgun":
		return NewShotgun(), nil

	default:
		return nil, errors.New("Wrong gun type!")
	}
}

func main() {
	pistol, _ := New("Pistol")
	shotgun, _ := New("Shotgun")

	fmt.Println(pistol.getType())
	fmt.Println(shotgun.getType())
}
