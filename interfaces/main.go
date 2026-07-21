package main

import "fmt"

// An interface is a set of method signature, it defines what a type does, rather than how it does it. it has a close relationship with
// struct, and this combination allows Go to have the ability to mimic polymorphism in OOP even though it is not object-oriented.

type Animal interface {
	sound() string
	move()
}

type Dog struct {
	name string
	age  int8
}

type Ostrich struct {
	name string
	age  int8
}

// Methods
func (dog Dog) sound() string {
	return "barks"
}

func (dog Dog) move() {
	println("walk on 4 legs")
}

func (ostrich Ostrich) sound() string {
	return "chirps"
}

func (ostrich Ostrich) move() {
	println("Flightless, runs on 2 legs")
}

func checkAnimal(animal Animal) {
	if animal.sound() == "barks" {
		fmt.Println("this is a dog")
	} else {
		fmt.Println("this is an ostrich")
	}
}

// More tricky examples 
// Measurable interface 
type Measurable interface {
	Value() float64
}

// Temperature is a custom type based on float64 
type Temperature float64

// Value receiver:  reads the number, doesn't need to touch anything
func (t Temperature) Value() float64 {
	return float64(t)
}

// Room is a struct
type Room struct {
	Width  float64
	Height float64
}

// Pointer receiver
func (r *Room) Value() float64 {
	return r.Width * r.Height
}

func main() {
	dog1 := Dog{name: "Edge", age: 2}
	ostrich1 := Ostrich{name: "Colo", age: 5}

	// Calling the method individually
	fmt.Println(dog1.sound())
	dog1.move()
	fmt.Println()

	fmt.Println(ostrich1.sound())
	ostrich1.move()

	checkAnimal(dog1)
	checkAnimal(ostrich1)

	// Animal interface example	 
	animals := []Animal{dog1, ostrich1}
	for _, animal := range animals {
		fmt.Println(animal.sound())
		animal.move()
	}

	// Measurable interface example
	var measure Measurable

	temparature := Temperature(98.6)
	measure = temparature
	fmt.Println("\nTemperature:", measure.Value()) // OK: Temperature (value) implements Measurable 

	room := &Room{Width: 10, Height: 20}
	measure = room
	fmt.Println("Room area:", measure.Value()) // OK: *Room (pointer) implements Measurable	

	// measure = Room{Width: 5, Height: 15} // FAILS: Room (value, not pointer) does NOT implement Measurable
	

}
