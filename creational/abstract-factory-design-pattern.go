package main

import "fmt"

/***
Abstract Factory is a creational design pattern that provides an interface for
creating families of related or dependent objects without specifying their concrete classes.
***/

// Product Interfaces
type Chair interface {
	SitOn()
}

// Concrete Products - Modern Style
type ModernChair struct{}

func (c *ModernChair) SitOn() {
	fmt.Println("Sitting on a modern chair.")
}

// Concrete Products - Victorian Style
type VictoriaChair struct{}

func (c *VictoriaChair) SitOn() {
	fmt.Println("Sitting on a victoria chair.")
}

// Abstract Factory Interface
type FurnitureFactory interface {
	CreateChair() Chair
}

// Concrete Factory - Modern
type ModernFurnitureFactory struct{}

func (f *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

// Concrete Factory - Victorian
type VictoriaFurnitureFactory struct{}

func (f *VictoriaFurnitureFactory) CreateChair() Chair {
	return &VictoriaChair{}
}

// Client code (Create Product)
func Client(factory FurnitureFactory) {
	chair := factory.CreateChair()
	chair.SitOn()
}

func main() {
	Client(&ModernFurnitureFactory{})
	Client(&VictoriaFurnitureFactory{})
}
