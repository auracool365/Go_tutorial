package main

import (
	"fmt"
)

// Defer is used to delay the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately,
// but the function call is not executed until the surrounding function returns. It is often used to ensure that resources are released or cleaned up
// properly at the end of a function e.g closing files, releasing locks, and closing databases, even if an error occurs or a function exits early.
// The deferred function calls are executed in LIFO (Last In, First Out) order, meaning that the last deferred function call is executed first.

// Function to demonstrate the use of defer with multiple return statements
func calculate(n int) int {
	defer fmt.Println("Deferred call: calculate function is about to return")
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n + calculate(n-1)
}

// Defer and Return Values: deferred functions execute after the return value has been computed but before the function actually returns.
func square() int {

	result := 4

	defer fmt.Println("Returning:", result)

	return result * result

}

func main() {
	defer fmt.Println("Last")
	fmt.Println("Printed First")
	fmt.Println()

	defer fmt.Println("Second to last")
	fmt.Println("Printed second")
	fmt.Println()

	defer fmt.Println("Third to last")
	fmt.Println("Printed third")
	fmt.Println()

	fmt.Println("count down")
	for i := 5; i >= 0; i-- {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	// Demonstrating defer with a function that has multiple return statements
	fmt.Println()
	fmt.Println("Demonstrating defer with a function that has multiple return statements")
	result := calculate(5)
	fmt.Println("Result:", result)

	// arguments are evaluated immediately: x evaluates immediately, only the execution of Println() is delayed.
	x := 10
	defer fmt.Println("defer x value", x)

	x = 50
	fmt.Println("Current value:", x)

	// Deferred anonymous functions: The function itself is deferred. Inside the function, y is looked up only when the deferred function actually runs.
	y := 20

	defer func() {
		fmt.Println("Inside anonymous defer:", y)
	}()
	fmt.Println("Outside defer:", y)

	y = 60

	// Return values: deferred Println() inside square() sees the local variable, not the returned expression.
	fmt.Println("Returned Expression Value:", square())

}
