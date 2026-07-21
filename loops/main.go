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
			fmt.Println("Skipped because of continue")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// Loops for enumeration
	// Looping through a string
	str := "Hello, World!"
	for i, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", i, char)
	}

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

	// Looping through an array
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
	fmt.Println()

	
	// While loop using for
	counter2 := 1
	for counter2 <= 10 {
		fmt.Println("Counter2: ", counter2)
		counter2++
	}

	// Infinite loop
	// for {
	// 	fmt.Println("This is an infinite loop")
	// }
	fmt.Println()


	// Nested loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
	fmt.Println()


	// Looping through a slice
	slice := []string{"apple", "peach", "pear", "watermelon", "guava"}
	for i, fruit := range slice {
		fmt.Printf("%d. %s \n", i, fruit)
	}
	fmt.Println()
	// Looping through a map
	fruitColors := map[string]string{
		"apple":      "red",
		"peach":      "orange",
		"pear":       "green",
		"watermelon": "green and red",
		"guava":      "green and yellow",
	}

	for fruit, color := range fruitColors {
		fmt.Printf("The color of %s is %s\n", fruit, color)
	}
	fmt.Println()


	// Looping through a channel
	fruitChannel := make(chan string, 5)
	fruitChannel <- "apple"
	fruitChannel <- "peach"
	fruitChannel <- "pear"
	fruitChannel <- "watermelon"
	fruitChannel <- "guava"
	close(fruitChannel)

	for fruit := range fruitChannel {
		fmt.Println(fruit)
	}
	fmt.Println()

	// Looping through a struct
	type Fruit struct {
		Name  string
		Color string
	}

	fruitsStruct := []Fruit{
		{"apple", "red"},
		{"peach", "orange"},
		{"pear", "green"},
		{"watermelon", "green and red"},
		{"guava", "green and yellow"},
	}
	fmt.Println(fruitsStruct)

	for i, fruit := range fruitsStruct {
		fmt.Printf("%d. %s is %s\n", i, fruit.Name, fruit.Color)
	}
	fmt.Println()

	// Looping through a pointer
	type FruitPointer struct {
		Name  string
		Color string
	}

	fruitsPointer := []*FruitPointer{
		{"apple", "red"},
		{"peach", "orange"},
		{"pear", "green"},
		{"watermelon", "green and red"},
		{"guava", "green and yellow"},
	}

	for i, fruit := range fruitsPointer {
		fmt.Printf("%d. %s is %s\n", i, fruit.Name, fruit.Color)
	}
	fmt.Println()

	// Looping through a slice of structs
	type FruitStruct struct {
		Name  string
		Color string
	}

	fruitsSliceStruct := []FruitStruct{
		{"apple", "red"},
		{"peach", "orange"},
		{"pear", "green"},
		{"watermelon", "green and red"},
		{"guava", "green and yellow"},
	}

	for i, fruit := range fruitsSliceStruct {
		fmt.Printf("%d. %s is %s\n", i, fruit.Name, fruit.Color)
	}
	fmt.Println()

	// Looping through a slice of pointers to structs
	type FruitPointerStruct struct {
		Name  string
		Color string
	}

	fruitsSlicePointerStruct := []*FruitPointerStruct{
		{"apple", "red"},
		{"peach", "orange"},
		{"pear", "green"},
		{"watermelon", "green and red"},
		{"guava", "green and yellow"},
	}

	for i, fruit := range fruitsSlicePointerStruct {
		fmt.Printf("%d. %s is %s\n", i, fruit.Name, fruit.Color)
	}
	fmt.Println()

	// Looping through a slice of slices
	sliceOfSlices := [][]string{
		{"apple", "peach", "pear"},
		{"watermelon", "guava"},
	}

	for i, slice := range sliceOfSlices {
		fmt.Printf("%d. %v\n", i, slice)
		for j, fruit := range slice {
			fmt.Printf("  %d.%d. %s\n", i, j, fruit)
		}
	}
	fmt.Println()

	// Looping through a slice of maps
	sliceOfMaps := []map[string]string{
		{"apple": "red", "peach": "orange"},
		{"pear": "green", "watermelon": "green and red"},
	}

	for i, fruitMap := range sliceOfMaps {
		fmt.Printf("%d. %v\n", i, fruitMap)
		for fruit, color := range fruitMap {
			fmt.Printf("  %d.%s is %s\n", i, fruit, color)
		}
	}
	fmt.Println()

	// Looping through a slice of channels
	sliceOfChannels := []chan string{
		make(chan string, 1),
		make(chan string, 1),
	}

	sliceOfChannels[0] <- "apple"
	sliceOfChannels[1] <- "peach"

	for i, fruitChannel := range sliceOfChannels {
		fmt.Printf("%d. %s\n", i, <-fruitChannel)
	}
}
