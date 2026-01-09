package main

import "fmt"

// Callbacks
func add(a, b int32) int32 { // if adjacent arguments are of the same type, the type can be written once
	return a + b
}

func multiply(a, b int32) int32 {
	return a * b
}

func sub(a, b int32) int32 {
	return a - b
}

// Higher-order function(HOF). Takes/return function as argument/result. ((callback_and_HOF args) callback_and_HOF return type)HOF return type {}
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

func main() {
	result := calculate(6, 4, add)
	fmt.Println("Addition:", result)

	result = calculate(4, 3, multiply)
	fmt.Println("Multiply:", result)

	result = calculate(6, 3, sub)
	fmt.Println("Subtraction:", result)

	arr := [4]string{"John", "Joseph", "Paul", "Precious"}
	greetings(arr[:], greetUser)
}
