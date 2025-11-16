package main

import (
	"fmt"
	"sort"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	// Values
	fmt.Println("Hi" + " Everyone") // concatenation
	fmt.Println("3.84 + 0.8 = ", 3.84+0.8)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// Variables
	// Static type declaration = variable keyword | variable name | data type
	var age int8 = 26
	fmt.Println(age)
	fmt.Println("I am", age, "years old")

	age = 27
	fmt.Println("I will be", age, "years old next month")

	const pi float32 = 3.142
	fmt.Println(pi)

	var userName string = "Cornelius"
	fmt.Println("My name is", userName)

	var isGoFun bool = true
	fmt.Println("Is golang fun?", isGoFun)

	city := "New York"
	fmt.Printf("The city is %s\n", city)

	// Multiple initialization
	var a, b, c int = 81, 82, 83
	fmt.Printf("a = %d\nb = %d\nc = %d\n", a, b, c)

	var x float32 = 13.6779
	var y int8 = 9
	// var z float32 = x + y // Error
	var z float32 = x + float32(y) // Fixed using type conversion
	fmt.Println(z)

	var num1 int8 = 7
	var num2 int8 = 3
	fmt.Println(num1 / num2)                   // Rounds up the answer
	fmt.Println(num1 % num2)                   // Returns the remainder
	fmt.Println(float32(num1) / float32(num2)) // this will give the proper division answer expected

	var myRune rune = 'a'
	fmt.Println("character 'a' in ASCII is ", myRune) // outputs 97

	// Variable declaration without initialization
	var myVar string
	fmt.Println(myVar) // outputs empty string " "

	var num int
	fmt.Println(num) // outputs 0(same as all other numeric types uint8 - uint64, int8 - int64, floats)

	// Data types can be ignored, the compiler will infer it
	var animal string
	animal = "Dog"
	fmt.Println(animal)

	animal = "Squirrel"
	fmt.Println(animal)

	var feline1, feline2 = "Cat", "Tiger"
	fmt.Println(feline1, feline2)

	// Even the var keyword can be ignored using this syntax
	// fruit string := "Apple" // Error!. Go does not allow specifying types for shorthand variable initialization
	fruit := "Apple"
	fmt.Println(fruit)

	name1, name2 := "Steve", "Wang"
	fmt.Println("Hello", name1)
	fmt.Println("Hello", name2)

	// Constants
	const speedOfLight int32 = 299_792_458 // m/s
	fmt.Println("The speed of light is", speedOfLight)
	// speed_of_light = 2767990 // Error! can't reassign a constant

	result1 := add(4, 5) // In this case shorthand is not recommended since I'd want to know the return type of the function
	fmt.Println(result1)

	var result2 int = add(4, 6) // Better, the return type is indicated
	fmt.Println(result2)

	// Check the type of a value(%T)
	var checkVar1, checkVar2, checkVar3 = "Hey", 3.142, 100
	fmt.Printf("checkVar1 is of type %T \n", checkVar1)
	fmt.Printf("checkVar2 is of type %T \n", checkVar2)
	fmt.Printf("checkVar3 is of type %T \n", checkVar3)

	// Using Blank Specifiers( _ ) to solve unused variables (Go does not allow unused variables)
	fruit1, fruit2, _ := "Dates", "Banana", "Watermelon" // The blank specifier prevents the unused variable error
	fmt.Printf("I love %s and %s😋 \n", fruit1, fruit2)

	var _ string = "Nature"

	// Intro to pointers
	num3 := 9
	num4 := num3

	// Pointer syntaxes
	var num3Ptr *int = &num3
	num4Ptr := &num4
	testPtr := &num3Ptr // pointer to pointer
	var zeroValuePtr *int

	num3++

	// *zeroValuePtr += 5 // Very bad(Segmentation fault similar to C++)

	fmt.Println("num3 =", num3)
	fmt.Println("num4 =", num4)

	fmt.Println("num3Ptr =", num3Ptr)
	fmt.Println("num3Ptr =", *num3Ptr)

	fmt.Println("num4Ptr =", num4Ptr)
	fmt.Println("num4Ptr =", *num4Ptr)

	fmt.Println("zeroValuePtr =", zeroValuePtr)
	// fmt.Println("zeroValuePtr =", *zeroValuePtr) // Very bad
	fmt.Println()

	*num3Ptr += 5
	**testPtr += 5
	*num4Ptr++

	fmt.Println("num3 =", num3)
	fmt.Println("num4 =", num4)

	fmt.Println("num3Ptr =", num3Ptr)
	fmt.Println("num3Ptr =", *num3Ptr)

	fmt.Println("num4Ptr =", num4Ptr)
	fmt.Println("num4Ptr =", *num4Ptr)

	fmt.Println("testPtr =", testPtr)     // address of num3Ptr
	fmt.Println("*testPtr =", *testPtr)   // address that num3Ptr points to
	fmt.Println("**testPtr =", **testPtr) // value stored at the address that num3Ptr is pointing to
	fmt.Println()

	// Pointers usefulness
	names := [3]string{"Alice", "Charlie", "Bob"}
	fmt.Println(names)
	secondName := names[1]     // Copy of the current second element(pass by value) will NOT change
	secondNamePtr := &names[1] // Address of the current second element(pass by reference) will change

	fmt.Println("The value of secondName is:", secondName)
	fmt.Println("secondNamePtr:", secondNamePtr)
	fmt.Println("*secondNamePtr:", *secondNamePtr)

	sort.Strings(names[:])
	fmt.Println(names)

	fmt.Println("The value of secondName is:", secondName) // Remains unchanged
	fmt.Println("secondNamePtr:", secondNamePtr)
	fmt.Println("*secondNamePtr:", *secondNamePtr) // Changed

}
