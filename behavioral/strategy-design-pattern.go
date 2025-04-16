package main

import "fmt"

/**
The Strategy design pattern defines a family of interchangeable algorithms,
encapsulated behind a common interface, allowing the behavior to change at runtime.
**/

type Animal interface {
	Sound()
}

type Cat struct{}

func (c *Cat) Sound() {
	fmt.Println("Cat sounds like Mew!!!")
}

type Dog struct{}

func (d *Dog) Sound() {
	fmt.Println("Dog sounds like Bark!!!")
}

func GenerateSound(animal Animal) {
	animal.Sound()
}

func main() {
	GenerateSound(&Dog{})
	GenerateSound(&Cat{})
}
