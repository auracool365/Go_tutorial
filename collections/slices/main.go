package main

import (
	"fmt"
	"reflect"
	"sort"
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

	mySlice = append(mySlice, 7, 8, 9, 10) // Adds multiple elements
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

	// Add the element from a slice using a loop
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

	// specifying range to copy from an array into a slice
	items1 := [5]string{"milk", "chocolate", "biscuits", "coke", "kilishi"}

	items2 := items1[1:3:3] // low, high, max
	items2 = append(items2, "bread")

	items3 := items2

	fmt.Println("items1:", items1)
	fmt.Println("items1 length:", len(items1), "capacity:", cap(items1))

	fmt.Println("items2:", items2)
	fmt.Println("items2 length:", len(items2), "capacity:", cap(items2))

	fmt.Println("items3:", items3)
	fmt.Println("items3 length:", len(items3), "capacity:", cap(items3))

	fmt.Println("Equal:", reflect.DeepEqual(items2, items3)) // true
	fmt.Println("Equal:", reflect.DeepEqual(items2, items1)) // false
	fmt.Println()

	// copy(destination, source)
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8}
	fmt.Println("nums1:", nums1)
	fmt.Println("nums1 length:", len(nums1), "capacity:", cap(nums1))

	fmt.Println("nums2:", nums2)
	fmt.Println("nums2 length:", len(nums2), "capacity:", cap(nums2))
	fmt.Println()

	copy(nums1, nums2) // output:[6 7 8 4 5] The source elements will replace elements in the destination starting from the beginning
	fmt.Println("nums1:", nums1)
	fmt.Println("nums1 length:", len(nums1), "capacity:", cap(nums1))

	// sort()
	fruits := []string{"Dates", "Watermelon", "Apple", "Banana", "Orange"}
	fmt.Println("fruits(unsorted):", fruits)

	sort.Strings(fruits)

	fmt.Println("fruits(sorted):", fruits)

}
