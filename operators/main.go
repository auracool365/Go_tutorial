package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operators")

	// Arithmetic operators
	var1, var2 := 250, 50
	sum := var1 + var2
	difference := var1 - var2
	product := var1 * var2
	remainder := var1 % var2

	fmt.Printf("var1 + var2 = %d\n", sum)
	fmt.Printf("var1 - var2 = %d\n", difference)
	fmt.Printf("var2 * var2 = %d\n", product)
	fmt.Printf("var2 mod var2 = %d\n", remainder)
	fmt.Printf("5 mod 2 = %d\n", 5%2)

	// Increment and decrement
	var3 := 9
	fmt.Println(var3)

	var3++
	fmt.Println(var3)

	var3 += 90
	fmt.Println(var3)

	// Concatenation

	text1 := "Hello"
	text2 := "World"
	message := text1 + " " + text2
	fmt.Println(message)

}
