package main

import "fmt"

// function composition is when multiple functions are combined to create a new one
func getName(userName string) string {
	return userName
}

func getMessage(message string) string {
	return message
}

// greetName() composes getName() and getMessage(), and returns a function. It has no idea how the name and message are produced
func greetName(getNameFunc func(string) string, getMessageFunc func(string) string) func(userName string, greetMessage string) string {

	// This returned function is the composed function
	return func(userName, greetMessage string) string {

		// Get the function parameters
		name := getNameFunc(userName)
		message := getMessageFunc(greetMessage)

		// Final composed result
		result := message + " " + name
		return result
	}
}

func main() {
	greetings := greetName(getName, getMessage)
	fmt.Println(greetings("Cornelius", "Hey"))
	fmt.Println(greetings("Tomisin", "Hello"))

}
