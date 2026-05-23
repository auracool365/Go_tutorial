package main

import (
	"errors"
	"fmt"
)

/* // Go does not use exceptions, to check for failures, functions simply return an error just as a normal return value.

func test() (valueType, error) {
	return valueType, error
}

Usage:
value, err := test()

if err != nil {
	failure
} else {
	success
} */

// using fmt.Errorf(): more flexible, allows formatting
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by %.2f", a, b)
	}
	return a / b, nil
}

// Using errors.New(): for simple static error messages
func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func greetUser(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	greeting := "Hello " + name
	return greeting, nil
}

// Custom error types
// 1. Define the custom error struct, adding extra fields like 'Field' allows callers to programmatically inspect what went wrong.
type ValidationError struct {
	Field   string
	Message string
}

// 2. Implement the Error() to satisfy the error interface
func (err *ValidationError) Error() string {
	return fmt.Sprintf("Invalid %s: %s", err.Field, err.Message)
}

// custom usages
func registerUser(username string) error {
	if username == "" {
		// 3. Return a pointer to the custom error struct
		return &ValidationError{
			Field:   "Username",
			Message: "cannot be empty",
		}
	}
	return nil
}

func checkAge(age int8) error {
	if age < 0 {
		return &ValidationError{
			Field: "UserAge",
			Message: "A person age can not be a negative number",
		}
	}
	return nil
}

func main() {
	// fmt.Errorf()
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// errors.New()
	result, err = divide2(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = divide2(10, 0)
	if err != nil {
		fmt.Println("shit:", err)
	} else {
		fmt.Println(result)
	}

	greeting, err := greetUser("Cornelius")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}

	greeting, err = greetUser("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}

	// Custom error types
	// err = registerUser("Steve") // no output
	err = registerUser("")

	if err != nil {
		fmt.Println("Caught an error:", err)
		// 4. Using errors.As to check for and access custom error fields
		var errorDetail *ValidationError
		if errors.As(err, &errorDetail) {
			fmt.Printf("Error details Access Field: %s, Msg: %s\n", errorDetail.Field, errorDetail.Message)
		}
	}

	//err = checkAge(5) // No output
	err = checkAge(-5)

	if err != nil {
		fmt.Println("Caught an error:", err)
		var errorDetail *ValidationError 
		if errors.As(err, &errorDetail) {
			fmt.Printf("Field: %s\nMessage: %s\n", errorDetail.Field, errorDetail.Message)
		}
	}

}
