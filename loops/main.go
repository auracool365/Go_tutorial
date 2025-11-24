package main

import "fmt"

func main() {
	const count int = 5
	i := 0

	// Without index
	for i < count {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println() // to add a new line

	i = 0
	for range count {
		fmt.Printf("%d. Hello\n", i)
		i++
	}
	fmt.Println()

	for character := 'a'; character <= 'z'; character++ {
		fmt.Printf("%c ", character)
	}

	// With index usage
	for i := 1; i <= count; i++ {
		fmt.Printf("%d. Hey \n", i)
	}
	fmt.Println()

	// Count down
	for i = 10; i >= 0; i-- {
		fmt.Println(i)
	}
	fmt.Println()

	// break statement
	for i := 0; i <= count; i++ {
		if i == 3 {
			fmt.Println("Stopped because of break!")
			break
		}
		fmt.Println(i)
	}

	// continue statement
	for i := 0; i <= count; i++ {
		if i == 3 {
			fmt.Println("❌Skipped because of continue")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// Loops for enumeration
	name := "Cornelius"
	for i, character := range name { // (i) get the index from range, character gets the values from the string
		fmt.Printf("%d. %c \n", i, character)
	}
	fmt.Println()

	for _, character := range name { // use underscore when the index is will not be used.
		fmt.Printf("%c \n", character)
	}
	fmt.Println()

	for i := range name {
		fmt.Println("index:", i)
	}

	fruits := [...]string{"apple", "peach", "pear", "watermelon", "guava"}
	for i, fruit := range fruits {
		fmt.Printf("%d. %s \n", i, fruit)
	}
	fmt.Println()

	for i, fruit := range fruits {
		switch fruit {
		case "watermelon":
			fmt.Printf("My favorite fruit is %s at index %d\n", fruit, i)
		case "apple":
			fmt.Println("I like the fruit at index", i, "also")
		case "peach":
			fmt.Println("Never had the fruit at index", i)
		case "pear":
			fmt.Println("Learning to eat the fruit at index", i)
		case "guava":
			fmt.Println("The fruit at index", i, "is okay")
		default:
			fmt.Println("No comment on this fruit")
		}
	}
	fmt.Println(fruits)
	fmt.Println()

	// Label: allows execution to jump to a different point for flexibility
	counter := 0
target: // Label
	fmt.Println("Counter", counter)
	counter++
	if counter < 10 {
		goto target
	}
}
