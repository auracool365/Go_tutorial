package main

import (
	"errors"
	"fmt"
)

func test(a int, b int) (int, int, error) { // Returns 3 things,
	var err error
	if b == 0 {
		err = errors.New("can not divide by zero")
		return 0, 0, err
	}
	return a + b, a / b, err // returns sum, quotient, error(if any)
}

func main() {
	age := 18
	// if
	if age < 18 {
		fmt.Println("You are under age")
	}

	// if and else
	if age < 18 {
		fmt.Println("You are a minor")
	} else {
		fmt.Println("You are an adult")
	}

	// if, else if, else
	if age > 17 && age < 20 {
		fmt.Println("You are in your late teens")
	} else if age > 20 && age < 30 {
		fmt.Println("You are a young adult")
	} else if age >= 30 && age < 50 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a senior")
	}

	var sum, quotient, errorResult = test(20, 5)

	if errorResult != nil {
		fmt.Println(errorResult.Error())
	} else {
		fmt.Printf("The sum from test() is %d \nThe quotient from test() is %d\n", sum, quotient)
	}

	// Switch statement
	switch {
	case errorResult != nil:
		fmt.Println(errorResult.Error())
	case sum > quotient:
		fmt.Println("Sum", sum, "is greater than quotient", quotient)
	case sum == quotient:
		fmt.Println("Sum", sum, "is equal to quotient", quotient)
	case sum < quotient:
		fmt.Println("Sum", sum, "is less than quotient", quotient)
	default:
		fmt.Println(sum, quotient)
	}

	// Switch statements can also have an expression
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	// The variable can also be created as part of the expression
	switch dayNum := 1; dayNum {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}
