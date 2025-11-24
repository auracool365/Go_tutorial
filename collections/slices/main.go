package main

import (
	"fmt"
	"time"
)

func timeMeasurement(slice []int, n int) time.Duration {
	timeVar := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(timeVar)
}

func main() {
	// Slices are like array but with more flexibility and dynamic usage
	var mySlice = []int16{1, 2, 3, 4, 5} // No need to specify the size
	fmt.Println(mySlice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 6) // adds a new element to the slice
	fmt.Println(mySlice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 7, 8, 9, 10)
	fmt.Println(mySlice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(mySlice), cap(mySlice))

	// creating a slice from other slices using the spread operator
	var mySlice2 = []int16{11, 12, 13, 14, 15}
	fmt.Println(mySlice2)

	mySlice = append(mySlice, mySlice2...)
	fmt.Println(mySlice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(mySlice), cap(mySlice))

	// Using make(type, size, capacity) for specifics. This is good if the size and capacity needed is known
	var mySlice3 = make([]int16, 5, 10)
	mySlice3[0] = 71
	mySlice3[1] = 72
	mySlice3[2] = 73
	mySlice3[3] = 74
	mySlice3[4] = 75
	fmt.Println(mySlice3)

	// Add the element from an array using a loop
	var mySlice4 []int16

	arr := [5]int16{1, 2, 3, 4, 5}
	for i := range len(arr) {
		mySlice4 = append(mySlice4, arr[i]+2) // adds 2 to each element before appending to the slice
	}
	fmt.Println(mySlice4)

	var mySlice5 []int16
	mySlice5 = append(mySlice5, arr[:]...) // [:] expands the array, ... appends its element to the slice
	fmt.Println(mySlice5)

	// Benchmark for specifying capacity vs not specifying capacity
	x := 1000000
	var sliceA []int                                                           // The compiler has to determine the capacity each time on its own
	sliceB := make([]int, 0, x)                                                // capacity is predetermined, the compiler doesn't have to do that
	fmt.Printf("Time without preallocation: %s\n", timeMeasurement(sliceA, x)) // slower
	fmt.Printf("Time with preallocation: %s\n", timeMeasurement(sliceB, x))    // much faster during my tests

	namesSlice := [4]string{"Sam", "John", "Paul", "Tom"}
	for index, element := range namesSlice {
		fmt.Printf("%d. %s\n", index, element)
	}

}
