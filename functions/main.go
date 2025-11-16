package main

import (
	"errors"
	"fmt"
	"math"
)

// Takes no parameter, receives no argument, does not return value
func greet() {
	fmt.Println("Hello Golang!")
}

// Takes parameter(s), has no return value
func greetUser(userName string) {
	fmt.Println("Hi", userName)
}

func add(a int32, b int32) {
	fmt.Println(a + b)
}

func findEven(num int) (int, string) {
	if num%2 == 0 {
		return num, "is even"
	}
	return num, "is odd"
}

// Has return value of type int32
func multiply(a int32, b int32) int32 {
	return a * b
}

// Return multiple values(subtraction, modulus)
func test(x int32, y int32) (int32, int32) {
	var a = x - y
	var b = x % y
	return a, b
}

// Adding the error type (division, power, error)
func divide(a int32, b int32) (int32, int32, error) {
	var err error
	if b == 0 {
		err = errors.New("can not divide by zero")
		return 0, 0, err
	}
	return a / b, int32(math.Pow(float64(a), float64(b))), err
}

// Variadic functions (can receive any number of arguments of the specified type)
func sum(nums ...int32) {
	fmt.Print(nums, " ") // The arguments will be printed as an array
	var total int32 = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("=", total)
}

// Recursion
func factorial(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	greet()
	greetUser("Cornelius")

	fmt.Println(multiply(2, 3))

	var x int32 = 20
	var y int32 = 60
	fmt.Println(multiply(x, y))

	fmt.Println("\nReturning multiple values (subtraction, modulus)")
	var difference, remainder int32 = test(15, 2)
	fmt.Println(difference, remainder)
	fmt.Printf("subtraction is %d and modulus is %d\n", difference, remainder) // C style printf for easier formatting
	// fmt.Printf("subtraction is %v and modulus is %v\n", difference, remainder) // %v is universal, generic for all data types

	add(x, y)

	fmt.Println("\nReturning multiple values (division, power, error)")
	var quotient, power, errorResult = divide(20, 5)

	if errorResult != nil {
		fmt.Println(errorResult.Error())
	} else {
		fmt.Println(quotient, power)
	}

	fmt.Println("\nVariadic function")
	sum(6, 67)

	sum(5, 10, 59)

	sum(1, 2, 3, 4, 5)

	arr := []int32{1, 2, 3, 4, 5, 6}
	sum(arr...)

	fmt.Println("\nRecursion")
	result := factorial(5)
	fmt.Println(result)

	fmt.Println(findEven(10))
	fmt.Println(findEven(11))
	fmt.Println(findEven(12))
	fmt.Println(findEven(13))
	fmt.Println(findEven(14))
	fmt.Println(findEven(15))

}
