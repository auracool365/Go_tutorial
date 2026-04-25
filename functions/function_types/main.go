package main

import "fmt"

// Function types; using functions as data types
func multiply(num int32) int32 {
	return num * 5
}

func add(num int32) int32 {
	return num + 10
}

func main() {
	arr := []int32{81, 20, 33, 74, 5}

	// Declaring a variable that holds a function. The type `func(int32) int32` means this variable can store any function
	// that accepts an int32 and returns an int32. The variable type is a function type that matches the defined signature.
	var calcFunc func(int32) int32

	for _, num := range arr {
		if num > 50 {
			// assigning the add() to the variable just like assigning a value to any regular variable.
			calcFunc = add
			addResult := calcFunc(num) // Invoking the function through the variable. Holds reference to add().
			fmt.Println("add: ", addResult)
		} else {
			// reassigning the variable to another function multiply(), swapping behavior at runtime without changing the call site.
			calcFunc = multiply
			multiplyResult := calcFunc(num)
			fmt.Println("multiply: ", multiplyResult)
		}
	}

}
