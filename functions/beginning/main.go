package main

import (
	"errors"
	"fmt"
	"math"
)

// Takes no parameter, receives no argument, no return value
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

// Has return value of type int32
func multiply(a int32, b int32) int32 {
	return a * b
}

// Returns int and string
func findEven(num int) (int, string) {
	if num%2 == 0 {
		return num, "is even"
	}
	return num, "is odd"
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

func calculateTax(productPrice float32) (float32, bool, float32) {
	taxRate := float32(10.0 / 100.0) // 10%

	// only tax product that cost 200 and above
	if productPrice >= 200 {
		tax := productPrice * taxRate
		totalAmount := productPrice + tax
		return tax, true, totalAmount
	}
	return 0, false, productPrice
}

// Variadic functions (can receive any number of arguments of the specified type)
func sum(nums ...int32) {
	fmt.Print(nums, " ") // The arguments will be printed as a slice
	var total int32 = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("=", total)
}

func greetMultipleNames(names ...string) {
	for index, name := range names {
		if len(name) < 5 {
			fmt.Printf("%d. Hello %s\n", index, name)
		} else {
			fmt.Printf("%d. Hi %s\n", index, name)
		}
	}
}

func printProfessionalAthlete(sport string, athletes ...string) {
	if len(athletes) == 0 {
		fmt.Println("Sport:", sport, "> Athlete: (No athletes)", athletes) // prints empty slice
		return
	}
	fmt.Println("Sport:", sport, "> Athlete:", athletes) // Prints the athlete as a slice
	for _, athlete := range athletes {
		fmt.Println("Sport:", sport, "> Athlete:", athlete)
	}

}

// Recursion
func factorial(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Pointers as parameters(pass by reference)
// pass by value(the limitation)
func increment1(num int) {
	num++
}

func swap1(text1 string, text2 string) {
	text1, text2 = text2, text1
}

// pass by reference(The solution)
func increment2(num *int) {
	*num++
}

func swap2(text1 *string, text2 *string) {
	*text1, *text2 = *text2, *text1
}

// Defer
func deferTest() {
	fmt.Println("function Started:")
	defer greetUser("John Doe")
	greetUser("Steve")
	fmt.Println("Function Finished.")
}

func main() {
	defer fmt.Println("defer keyword test.") // push unto the bottom of the execution

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

	products := map[string]float32{"iPhone": 1299, "socks": 12}
	for product, amount := range products {
		taxAmount, taxable, totalAmount := calculateTax(amount)

		if taxable {
			fmt.Printf("Product = %s, Tax = %.2f, Total price = %.2f\n", product, taxAmount, totalAmount)
		} else {
			fmt.Printf("Product = %s, Tax = %.2f, Total price = %.2f\n", product, taxAmount, totalAmount)
		}
	}

	fmt.Println("\nVariadic functions")
	sum(6, 67)

	sum(5, 10, 59)

	sum(1, 2, 3, 4, 5)

	arr := []int32{1, 2, 3, 4, 5, 6}
	sum(arr...)

	greetMultipleNames("Sam", "Cornelius", "Steve", "John", "Philip")

	namesArr := []string{"Jannick Sinner", "Novak Djokovic"}
	printProfessionalAthlete("Football", "Cristiano Ronaldo", "Lionel Messi", "Zinadine Zidane")
	printProfessionalAthlete("Tennis", namesArr...)
	printProfessionalAthlete("LineBall")

	fmt.Println("\nRecursion")
	result := factorial(5)
	fmt.Println(result)

	fmt.Println(findEven(10))
	fmt.Println(findEven(11))
	fmt.Println(findEven(12))
	fmt.Println(findEven(13))
	fmt.Println(findEven(14))
	fmt.Println(findEven(15))

	fmt.Println(100 - 99 + 1)

	fmt.Println("\nPass by value")
	num1 := 4
	fmt.Println("Before increment:", num1)
	increment1(num1)
	fmt.Println("After increment:", num1)

	fruit1, fruit2 := "Apple", "Orange"
	fmt.Println("Before swap:", fruit1, fruit2)
	swap1(fruit1, fruit2)
	fmt.Println("After swap:", fruit1, fruit2)

	fmt.Println("\nPass by reference")
	num2 := 9
	fmt.Println("Before increment:", num2)
	increment2(&num2)
	fmt.Println("After increment:", num2)

	fruit3, fruit4 := "Watermelon", "Dates"
	fmt.Println("Before swap:", fruit3, fruit4)
	swap2(&fruit3, &fruit4)
	fmt.Println("After swap:", fruit3, fruit4)

	fmt.Print("\nDefer keyword \n\n")
	defer fmt.Println("Cornelius ")
	fmt.Println("Hello ")

	deferTest()

}
