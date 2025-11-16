package main

import (
	"fmt"
	"math/rand"
)

func greet(name string) string {
	return "Hello " + name
}

func introduce(name string) {
	fmt.Println("Hi, my name is", name)
}

func main() {
	userName := "Cornelius"
	const message string = "Hi"
	fmt.Println(message, userName)

	userName = "Steve"
	fmt.Println(greet(userName))
	introduce(userName)

	var nums = [...]int16{78, 86, 71, 87, 94, 73, 82}
	for i := range nums {
		if nums[i]&1 == 1 {
			fmt.Println(nums[i], "is an odd number!")
		} else {
			fmt.Println(nums[i], "is an even number!")
		}
	}

	var randomValue uint64 = rand.Uint64()
	fmt.Println("The random value is ", randomValue)
}
