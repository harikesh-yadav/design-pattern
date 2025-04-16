package main

import "fmt"

/**
The Factory Method pattern is a creational design pattern that provides an interface
for creating objects in a superclass but allows subclasses to alter the type of objects that will be created.
**/

// Step 1: Define the Product interface
type Shape interface {
	Draw()
}

// Step 2: Concrete Products
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a Circle!!!")
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing a Rectangle!!!")
}

// Step 3: Factory Method
func GetShape(shapeType string) Shape {
	switch shapeType {
	case "CIRCLE":
		return &Circle{}
	case "RECTANGLE":
		return &Rectangle{}
	}
	return nil
}

func main() {
	circle := GetShape("CIRCLE")
	circle.Draw()
	rectangle := GetShape("RECTANGLE")
	rectangle.Draw()
}
