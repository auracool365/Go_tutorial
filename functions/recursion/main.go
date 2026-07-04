package main

import (
	"fmt"
	"strings"
)

// A recursive function is a function that calls itself in order to solve a problem. It typically has a base case that stops
// the recursion and a recursive case that breaks the problem into smaller sub-problems.

// Example 1: increment function
// Using loop
func incrementLoop(num int) int {
	for range 5 {
		num++
	}
	return num
}

// Using recursion
func increment(num int) int {
	// Base case: stop after incrementing 5 times
	if num >= 5 {
		return num
	}
	// Recursive case
	return increment(num + 1)
}

// Example 2: Factorial
// Using loop
func factorialLoop(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

// Using recursion
func factorial(num int) int {
	// Base case
	if num == 0 {
		return 1
	}
	// Recursive case
	return num * factorial(num-1)
}

// Example 3: join strings in a slice or array
// Using loop
func joinStringsLoop(words []string) string {
	var result strings.Builder
	for _, str := range words {
		result.WriteString(str)
	}
	return result.String()
}

// Using recursion
func joinStrings(words []string) string {
	// Base case: if the slice is empty, return an empty string
	if len(words) == 0 {
		return ""
	}
	// Recursive case: take the first string and concatenate it with the result of joining the rest of the slice
	return words[0] + joinStrings(words[1:])
}

// Example 4: multiply by repeated addition
// Using loop
func multiplyLoop(a, b int) int {
	result := 0
	for range b {
		result += a
	}
	return result
}

// Using recursion
func multiply(a, b int) int {
	// Base cases
	if a == 0 || b == 0 {
		return 0
	}

	if b == 1 {
		return a
	}
	// Handle negative numbers(more comprehensive than the dart example, which handled only positive numbers)
	if a < 0 && b < 0 {
		a = -a
		b = -b
	}
	if b < 0 {
		return -multiply(a, -b)
	}
	if a < 0 {
		return -multiply(-a, b)
	}
	// Recursive case: add a to the result of multiplying a by (b-1)
	return a + multiply(a, b-1)
}

// Example 5: Tower of Hanoi
// The Tower of Hanoi is a classic problem in which the goal is to move a stack of disks from one peg to another, following these specific rules:
// 1. Only one disk can be moved at a time.
// 2. Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack or on an empty peg.
// 3. No disk may be placed on top of a smaller disk.

// Using loop
func towerOfHanoiLoop(numberOfDisks int, source, target, auxiliary string) {
	// This problem is typically solved using recursion, but it can be implemented using an iterative approach with a stack.
	type Move struct {
		numberOfDisks int
		source   string
		target   string
		auxiliary string
	}

	stack := []Move{{numberOfDisks, source, target, auxiliary}}

	for len(stack) > 0 {
		move := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if move.numberOfDisks == 1 {
			fmt.Printf("Move disk 1 from %s to %s\n", move.source, move.target)
		} else {
			stack = append(stack, Move{move.numberOfDisks - 1, move.auxiliary, move.target, move.source})
			fmt.Printf("Move disk %d from %s to %s\n", move.numberOfDisks, move.source, move.target)
			stack = append(stack, Move{move.numberOfDisks - 1, move.source, move.auxiliary, move.target})
		}
	}
}

// Using recursion
func towerOfHanoi(numberOfDisks int, source, target, auxiliary string) {
	// Base case: if there is only one disk, move it from source to target
	if numberOfDisks == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", source, target)
		return
	}
	// Recursive case:
	// Move numberOfDisks-1 disks from source to auxiliary
	towerOfHanoi(numberOfDisks-1, source, auxiliary, target)

	// Move the nth disk from source to target
	fmt.Printf("Move disk %d from %s to %s\n", numberOfDisks, source, target)

	// Move the numberOfDisks-1 disks from auxiliary to target
	towerOfHanoi(numberOfDisks-1, auxiliary, target, source)
}

func main() {
	// increment function
	fmt.Println("Increment using loop:", incrementLoop(0))
	fmt.Println("Increment using recursion:", increment(0))

	// factorial function
	fmt.Println("Factorial using loop:", factorialLoop(5))
	fmt.Println("Factorial using recursion:", factorial(5))

	// join strings function
	words := []string{"Hello, ", "world!", " How are you?"}

	fmt.Println("Join strings using loop:", joinStringsLoop(words))
	fmt.Println("Join strings using recursion:", joinStrings(words))

	// multiply function
	fmt.Println("Multiply using loop:", multiplyLoop(5, 3))
	fmt.Println("Multiply using recursion:", multiply(5, 3))
	fmt.Println("Multiply with negative numbers using recursion:", multiply(-5, -3))
	fmt.Println("Multiply with negative numbers using recursion:", multiply(-5, 3))
	fmt.Println("Multiply with negative numbers using recursion:", multiply(5, -3))

	// tower of hanoi function
	fmt.Println("\nTower of Hanoi with 3 disks(Loop):")
	towerOfHanoiLoop(3, "A", "C", "B")
	fmt.Println("\nTower of Hanoi with 4 disks(Recursion):")
	towerOfHanoi(4, "A", "C", "B")
}
