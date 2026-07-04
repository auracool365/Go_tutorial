package main

import (
	"fmt"
	"time"
)

/* Callbacks: A callback is a function that is passed as an argument to another function and is executed after some operation has been
completed. It allows for asynchronous programming and can be used to handle events or perform actions after a certain task is done.

Higher-order functions (HOF): A higher-order function is a function that takes one or more functions as arguments and/or returns a
function as its result. HOFs are a powerful tool in functional programming and can be used to create more flexible and reusable code.
*/

func add(a, b int32) int32 { // if adjacent arguments are of the same type, the type can be written once
	return a + b
}

func multiply(a, b int32) int32 {
	return a * b
}

func sub(a, b int32) int32 {
	return a - b
}

// Higher-order function (HOF)
func calculate(a int32, b int32, callBack func(a int32, b int32) int32) int32 {
	return callBack(a, b)
}

// Callback
func greetUser(username string) string {
	if len(username) > 5 {
		return "Hi " + username
	}
	return "Hello " + username
}

// HOF. Takes function as argument, but doesn't return as result
func greetings(names []string, callback func(string) string) {
	for index, username := range names {
		fmt.Printf("%d. ", index)
		fmt.Println(callback(username))
	}
}

// HOF. Takes function as argument, and returns a function as result
func createGreeter(greeting string) func(string) string {
	return func(name string) string {
		return greeting + " " + name
	}
}


func main() {
	result := calculate(6, 4, add)
	fmt.Println("Addition:", result)

	result = calculate(4, 3, multiply)
	fmt.Println("Multiply:", result)

	result = calculate(6, 3, sub)
	fmt.Println("Subtraction:", result)

	arr := [4]string{"John", "Joseph", "Paul", "Precious"}
	greetings(arr[:], greetUser)

	greeter := createGreeter("Welcome")
	fmt.Println(greeter("Alice"))
	fmt.Println(greeter("Bob"))
	fmt.Println(greeter("Charlie"))
	// Using the same HOF with a different callback function
	farewellGreeter := createGreeter("Goodbye")
	fmt.Println(farewellGreeter("Alice"))
	fmt.Println(farewellGreeter("Bob"))
	fmt.Println(farewellGreeter("Charlie"))

	// using the createGreeter HOF, asynchronously greet names in a array welcome, then after 5 seconds, greet the same names goodbye
	// This is just a simulation of asynchronous behavior using goroutines and sleep, in a real application, you might use channels 
	// will be used to coordinate the execution of the greetings and farewells more effectively.
	fmt.Println("Asynchronous Program")
	go func() {
		greetings(arr[:], createGreeter("Welcome"))
	}()
	// Simulate a delay before saying goodbye
	// In a real application, you might use channels or other synchronization methods to coordinate this
	time.Sleep(1 * time.Second)
	fmt.Println("Waiting for 5 seconds before saying goodbye...")
	time.Sleep(5 * time.Second)
	greetings(arr[:], createGreeter("Goodbye"))
}
