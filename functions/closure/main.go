package main

import "fmt"

/* closure is a special type of anonymous function that references and retains access to variables from its surrounding scope,
even after the outer function has finished executing */

// counter() returns a function, not an int
func counter() func() int {

	// count is local to counter(). Normally, this would be destroyed when counter() finishes
	count := 0

	// returns an inner anonymous function (closure). This function "closes over" the variable `count`
	return func() int {

		// Each time this inner function is called, it modifies the SAME `count` variable in the outer scope
		count++

		// Return the updated value
		return count
	}
}

func main() {

	// counter() is executed ONCE here
	// It returns the inner function and assigns it to `result`
	// IMPORTANT: `count` is NOT destroyed — it is kept alive
	result := counter()

	// Each call invokes the SAME inner function
	// which still has access to the same `count`
	fmt.Println(result()) // 1
	fmt.Println(result()) // 2
	fmt.Println(result()) // 3
	fmt.Println(result()) // 4
}
