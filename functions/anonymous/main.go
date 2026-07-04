package main

import "fmt"

/* Anonymous functions: An anonymous function is a function that is defined without a name. It can be assigned to a variable, 
passed as an argument to another function, or immediately invoked. Anonymous functions are often used for short-lived operations or 
when a function is needed as a callback. 
*/

// passing as a parameter 
func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}

// immediately invoked function expression (IIFE)
var result = func(x, y int) int {
	return x + y
}(10, 20)

// closure returning an anonymous function 
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	// Example 1: Assigning an anonymous function to a variable
	greet := func() {
		fmt.Println("Hello from an anonymous function!")
	}
	greet() 

	greetWithName := func(name string) {
		if len(name) > 5 {
			fmt.Printf("Hello, %s!\n", name)
			return
		}
		fmt.Printf("Hi, %s!\n", name)
	}
	greetWithName("Cornelius")
	greetWithName("John")

	// IIFE
	calculation := func(x, y int) int {
		return x * y
	}(5, 4) 
	fmt.Println("Product of x and y:", calculation)

	// Example 2: Passing an anonymous function as an argument
	numbers := []int{1, 2, 3, 4, 5}
	sum := func(nums []int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}(numbers) 
	fmt.Println("Sum of numbers:", sum) 

	// Example 3: Using an anonymous function as a callback
	arr := []string{"Alice", "Bob", "Charlie"}
	for index, name := range arr {
		func(i int, n string) {
			fmt.Printf("%d	. Hello %s\n", i, n)
		}(index, name) 
	}

	// Example 4: Using an anonymous function with a higher-order function
	add := func(x, y int) int {
		return x + y
	}
	result := operate(5, 3, add)
	fmt.Println("addition:", result)

	multiply := func(x, y int) int {
		return x * y
	}
	result = operate(5, 3, multiply)
	fmt.Println("multiplication:", result)

	// Example 5: Using an anonymous function to create a closure
	count := counter()
	fmt.Println("Counter :", count())
	fmt.Println("Counter :", count())
	fmt.Println("Counter :", count())


	// 6. Storing anonymous functions in maps
	operations := map[string]func(int, int) int{
		"add": func(x, y int) int {
			return x + y
		},
		"subtract": func(x, y int) int {
			return x - y
		},
		"multiply": func(x, y int) int {
			return x * y
		},
		"divide": func(x, y int) int {
			if y == 0 {
				fmt.Println("Cannot divide by zero")
				return 0
			}
			return x / y
		},
	}

	fmt.Println("Addition from map:", operations["add"](10, 5))
	fmt.Println("Subtraction from map:", operations["subtract"](10, 5))
	fmt.Println("Multiplication from map:", operations["multiply"](10, 5))
	fmt.Println("Division from map:", operations["divide"](10, 5))


	emoji := "😀"
	fmt.Printf("Emoji: %s, Unicode: U+%04X", emoji, emoji)
	fmt.Println(len(emoji))

}