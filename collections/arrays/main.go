package main

import "fmt"

func main() {
	// Array collection in Go is similar to the C-style array. Stores fixed number of elements
	// Declaration
	var nums [4]int8
	fmt.Println(nums) // filled with [0 0 0 0]

	nums[2] = 85
	fmt.Println(nums) // filled with [0 0 85 0]

	nums[0] = 81
	nums[1] = 82
	nums[3] = 83
	fmt.Println(nums) // filled with [81 82 85 83]

	fmt.Println(nums[1:3]) // prints elements at index 1 and 3 [82 85]

	// Memory addresses
	fmt.Println(&nums) // &[81 82 85 83]
	fmt.Println(&nums[0])
	fmt.Println(&nums[1])
	fmt.Println(&nums[2])
	fmt.Println(&nums[3])

	// Using loops print individually rather than as a collection of values
	for i := range nums {
		fmt.Printf("%d. %d \n", i, nums[i])
	}

	// Array initialization
	var fruits = [3]string{"apple", "orange", "banana"}
	fmt.Println(fruits)

	for i := range len(fruits) {
		fmt.Println(fruits[i])
	}

	animals := [4]string{"dog", "cat", "bird"} // an empty space is added for the last element not added
	fmt.Println(animals)

	//arr := [3]int{1, 2, 3, 4} // Error. Out of bounds
	arr := [...]int8{81, 72, 93, 84, 75, 66, 87, 78, 89} // Dynamic size.
	fmt.Println(arr)

	for index, element := range arr {
		fmt.Printf("%v. %v \n", index, element)
	}

	names := [3]string{"Paul", "Steve", "Jack"}
	fmt.Println(names)

	// In Go, array type is a combination of its fixed length and elements type
	// var test1 [4]string = names //Error: Cannot use 'names' (type [3]string) as the type [4]string
	// fmt.Println(test1)

	// This works. Since the length of test2 wasn't specified, the type(length and element type) of names is copied and given to test2
	test2 := names
	fmt.Println(test2)
	fmt.Println()
	sameness := names == test2 // true
	fmt.Println(sameness)

	// Pass by value
	fmt.Println("Pass by value")
	names[0] = "John"
	fmt.Println("names:", names)
	fmt.Println("test2:", test2)
	fmt.Println()

	sameness = names == test2 // still true even though they hold different values, because their type(length and element type) are the same
	fmt.Println(sameness)

	// To not create copies, a pointer can be used
	test3 := &names // test3 is a pointer to [3] string

	// Pass by pointer
	fmt.Println("Pass by pointer")
	names[1] = "Pharrell"
	fmt.Println("names:", names)
	fmt.Println("test3:", *test3)

	// sameness = names == test3 // Error: type of [3] string != *[3] string
	fmt.Println()

	// Iterating over an array
	felines := [...]string{"Lion", "Tiger", "Leopard", "Jaguar", "Cougar"}
	for i, element := range felines {
		fmt.Printf("%d. %s \n", i, element)
	}

}
