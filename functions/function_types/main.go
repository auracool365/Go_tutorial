package main

import "fmt"

// Function types; using functions as data types

func multiply(num int32) int32 {
	return num * 5
}

func add(num int32) int32 {
	return num + 10
}

func main() {
	arr := []int32{81, 20, 33, 74, 5}

	var calcFunc func(int32) int32

	for _, num := range arr {
		if num > 50 {
			calcFunc = add
			addResult := calcFunc(num)
			fmt.Println("add: ", addResult)
		} else {
			calcFunc = multiply
			multiplyResult := calcFunc(num)
			fmt.Println("multiply: ", multiplyResult)
		}
	}

}
