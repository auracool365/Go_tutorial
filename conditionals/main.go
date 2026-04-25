package main

import (
	"errors"
	"fmt"
	"strconv"
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

	// Nested if statements
	age = 20
	isCitizen := true
	hasPVC := true

	if isCitizen {
		if age >= 18 {
			if hasPVC {
				fmt.Println("You can proceed to vote")
			} else {
				fmt.Println("You must have a PVC to vote")
			}
		} else {
			fmt.Println("You must be 18 or older to vote")
		}
	} else {
		fmt.Println("You must be a citizen to vote")
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

	// More examples. Note: Go has no ternary operators
	weather := "🌞"
	if weather == "🌞" {
		fmt.Println("The weather is sunny")
	} else {
		fmt.Println("The weather is not sunny")
	}

	names := [5]string{"Sam", "John", "Bob", "Mary", "Jane"}
	// I will check later to know which of the loop syntax is most recommended
	for i, name := range names {
		if len(name) <= 3 {
			fmt.Printf("%d. Hello %s!\n", i, name)
		} else {
			fmt.Printf("%d. Hi %s!\n", i, name)
		}
	}

	for i := 0; i < len(names); i++ {
		if len(names[i]) <= 3 {
			fmt.Printf("%d. Hello %s!\n", i, names[i])
		} else {
			fmt.Printf("%d. Hi %s!\n", i, names[i])
		}
	}

	stringVar := "500"
	fmt.Println(stringVar)
	if convertToInt, err := strconv.Atoi(stringVar); err == nil {
		fmt.Println("Converted to int:", convertToInt)
	} else {
		fmt.Println("Error:", err)
	}
	fmt.Println()

	word := "Baby"
	for i, char := range word {
		switch char {
		case 'B', 'b':
			if char == 'b' {
				fmt.Printf("lowercase %c at index %d\n", char, i)
				break
			}
			fmt.Printf("uppercase %c at index %d\n", char, i)
		case 'a':
			fmt.Println("a at index", i)
		}
	}

	// Switch without comparison value
	for counter := range 10 {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}

	// Nested switch statements
	age = 25
	switch {
	case age < 18:
		fmt.Println("You are a minor")
	case age >= 18 && age < 30:
		fmt.Println("You are a young adult")
		switch {
		case age >= 18 && age < 21:
			fmt.Println("You are between 18 and 21")
		case age >= 21 && age < 25:
			fmt.Println("You are between 21 and 25")
		default:
			fmt.Println("You are between 25 and 30")
		}
	default:
		fmt.Println("You are an adult")
	}
}
