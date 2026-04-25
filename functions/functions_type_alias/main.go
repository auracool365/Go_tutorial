package main

import "fmt"

// Functions type alias:
// Assigning a name to a function signature. So, the parameter and result types are not specified every time the function type is
// used. It is a way to create a new type that represents a function signature, making the code more readable and maintainable.
// E.g, instead of writing func(int32) int32 everywhere, that signature gets a name, so,any function that takes an int32 and
// returns an int32 satisfies this type.

type aliasFuncName func(int32) int32

func multiply(num int32) int32 {
	return num * 5
}

func add(num int32) int32 {
	return num + 10
}

// Usage: using aliasFuncName as a parameter type instead of the verbose func(int32) int32 — this is the readability benefit.
func calculate(arr []int32, calcFunc aliasFuncName) {
	for _, num := range arr {
		result := calcFunc(num)
		fmt.Println(result)
	}
}

// return a new array
func calculateOnArray(arr []int32, calcFunc aliasFuncName) []int32 {
	result := make([]int32, len(arr))
	for i, num := range arr {
		result[i] = calcFunc(num)
	}
	return result
}

// Using CalcFunc as a return type, selecting and returning a function based on a condition.
func selectFunction(condition string) aliasFuncName {
	switch condition {
	case "add":
		return add
	case "multiply":
		return multiply
	default:
			fmt.Println("Invalid condition, returning nil")
			return nil
	}
}

func main() {
	arr := []int32{81, 20, 33, 74, 5}

	calculate(arr, add) // Add 10 to each element.
	calculate(arr, multiply) // Multiply each element by 5.

	result := calculateOnArray(arr, add) // Add 10 to each element, and return the new array.
	fmt.Println("Result of calculateOnArray with add:", result)

	result = calculateOnArray(arr, multiply) // Multiply each element by 5, and return the new array.
	fmt.Println("Result of calculateOnArray with multiply:", result)

	// Using selectFunction to get a function based on a condition and then invoking it.
	chosenFunc := selectFunction("add") // Adding 10 to 10 using the selected function
	if chosenFunc != nil {
		fmt.Println("Selected function result for 10:", chosenFunc(10))
	}

	chosenFunc = selectFunction("multiply") // Multiplying 10 by 5
	if chosenFunc != nil {
		fmt.Println("Selected function result for 10:", chosenFunc(10))
	}

	chosenFunc = selectFunction("Subtract") // Invalid condition, should return nil and print an error message.
	if chosenFunc != nil {
		fmt.Println("Selected function result for 10:", chosenFunc(10))
	}	
}
