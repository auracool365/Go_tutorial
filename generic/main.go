package main

import "fmt"

// Generics enable generic programming(Pun intended 😅) in Go

// without generics
func addInts(a, b int) int {
	return a + b
}

func addFloats(a, b float64) float64 {
	return a + b
}

// With generics
// 1. Number types allowed
type Number interface {
	int | float64
}

// Generic add function
func add[T Number](a, b T) T {
	return a + b
}

// Working with Slice
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// 2. with Structs
// Generic container: any makes the generic unrestricted, no operations is guaranteed, anything can be tried, but can't be assumed.
type Box[T any] struct {
	Value T
}

// Print stored value
func (b Box[T]) print() {
	fmt.Println(b.Value)
}

// 3. Generic stack
type Stack[T any] struct {
	items []T
}

// Push item
func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

// Pop item
func (s *Stack[T]) pop() T {

	last := len(s.items) - 1

	value := s.items[last]

	s.items = s.items[:last]

	return value
}

// Constraints
// Compare two values: any will not work, because not all types are comparable, e.g slices, map, function types. comparable allows Go to use force
/* func Equal[T any](a, b T) bool {
	return a == b
} */
func equal[T comparable](a, b T) bool {
	return a == b
}

// Custom constraints, specify the types allowed
// Numeric types allowed
type Numeric interface {
	int | int64 | float32 | float64
}

// Return larger value
func max[T Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic Interface
// Generic repository
type Repository[T any] interface {
	Save(item T)
}

// User model
type User struct {
	Name string
}

// User repository
type UserRepository struct{}

// Save user
func (u UserRepository) Save(user User) {
	fmt.Println("Saved:", user.Name)
}

func main() {
	// without generics
	fmt.Println(addInts(4, 5))
	fmt.Println(addFloats(4.5, 6.0))
	fmt.Println()

	// With generics
	fmt.Println(add(4, 6))
	fmt.Println(add(2.5, 7.5))
	fmt.Println()

	// Working with Slice
	numbers := []int{1, 2, 3}

	names := []string{"Ade", "Bola", "Mariam"}

	printSlice(numbers)

	printSlice(names)
	fmt.Println()

	// Structs
	intBox := Box[int]{
		Value: 100,
	}

	stringBox := Box[string]{
		Value: "Hello",
	}

	intBox.print()

	stringBox.print()

	// Stack
	var intStack Stack[int]

	intStack.push(10)
	intStack.push(20)
	intStack.push(30)

	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())

	var stringStack Stack[string]

	stringStack.push("Go")
	stringStack.push("Rust")

	fmt.Println(stringStack.pop())

	// Constraints
	fmt.Println(equal(10, 10))
	fmt.Println(equal("Go", "Cpp"))
	fmt.Println(equal(true, true))

	// Custom Constraints
	fmt.Println(max(20, 50))
	fmt.Println(max(3.14, 1.25))

	// Interface with generics
	var repo Repository[User]
	repo = UserRepository{}
	repo.Save(User{
		Name: "Cornelius",
	})
}
