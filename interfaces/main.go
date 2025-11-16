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

	//
	animals := []Animal{dog1, ostrich1}
	for _, animal := range animals {
		fmt.Println(animal.sound())
		animal.move()
	}

}
