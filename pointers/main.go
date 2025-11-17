package main

import "fmt"

// Pass by value
func increment1(arr [5]uint8) [5]uint8 {
	for i := range len(arr) {
		arr[i]++
	}
	fmt.Printf("Array address in increment_1() %p\n", &arr)
	return arr
}

// Pass by reference
func increment2(arr *[5]uint8) [5]uint8 {
	for i := range len(arr) {
		arr[i]++
	}
	fmt.Printf("Address in increment_2() %p\n", arr)
	return *arr
}

func changeName(userName *string) *string {
	*userName = "Johnny"

	return userName
}

func main() {
	var myPtr *int32   // Declared but not initialized, pointer points to nil
	fmt.Println(myPtr) // nil
	// fmt.Println(my_ptr) // Error! trying to read value address that doesn't exist, since my_ptr has not point to any address

	var myVar int32 = 10
	myPtr = &myVar
	fmt.Println("The value is", myVar, "at address", myPtr)

	// Dereferencing
	*myPtr = 20
	fmt.Println("The value is", *myPtr, "at address", myPtr)

	// Initialize a pointer with new() and take memory space to hold value
	var userName *string = new(string)
	fmt.Println("The value is", *userName, "at address", userName) // No value at the moment, but the pointer points to a legit address

	*userName = "Cornelius" // writes a value to that address
	fmt.Println("The value is", *userName, "at address", userName)

	// Initialize with a variable address
	isGoFun := true
	var isGoFunPtr *bool = &isGoFun
	fmt.Println("The value is", isGoFun, "at address", isGoFunPtr)

	// If a shorthand syntax variable is used, then it is better to also use short hand for a pointer, to avoid type error
	a := 3.142
	aPtr := &a
	fmt.Println("The value is", a, "at address", aPtr)

	// Reassigning pointers
	var x float32 = 20.001
	var xPtr *float32 = &x
	fmt.Println("The value is", x, "at address", xPtr)

	var y float32 = 30.002
	xPtr = &y
	fmt.Println("The value is", y, "at address", xPtr)
	fmt.Println()

	// Function with pointers
	// Copying (pass by value) leading to different address
	nums := [5]uint8{1, 2, 3, 4, 5}
	fmt.Printf("%d Array address in main() %p\n", nums, &nums)
	fmt.Println(increment1(nums)) // Incremented copies inside increment_1()
	fmt.Println(nums)             // The original elements remained unchanged
	fmt.Println()

	// pass by reference (no copies), same address
	fmt.Printf("%d Array address in main() %p\n", nums, &nums)
	fmt.Println(increment2(&nums)) // Accepts the memory address of the array and modifies the original elements of the array
	fmt.Println(nums)              // The original elements show the change

	name := "Tom"
	fmt.Println(name)
	changeName(&name)
	fmt.Println(name)

	name1 := "Ade"
	fmt.Println(name1)
	changeName(&name1)
	fmt.Println(name1)

}
