package main

import "fmt"

func main() {
	// Map like in other languages is used to hold data as key-value pairs
	// Syntax : map[key type] value type{key : value}
	var countryAndCity = map[string]string{"USA": "New York", "India": "Delhi", "Japan": "Tokyo", "China": "Shanghai"}
	fmt.Println(countryAndCity)
	fmt.Println(countryAndCity["USA"]) // get the values using the key

	numAndAlphabet := map[uint8]string{1: "a", 2: "b", 3: "c", 4: "d"}
	fmt.Println(numAndAlphabet)
	fmt.Println(numAndAlphabet[1])

	// Using make(map[key type] value type)
	sweetFruits := make(map[string]bool)

	sweetFruits["apple"] = true
	sweetFruits["banana"] = true
	sweetFruits["lime"] = false
	sweetFruits["bitter kola"] = false
	sweetFruits["dates"] = true
	sweetFruits["mango"] = true
	fmt.Println(sweetFruits)
	fmt.Println(sweetFruits["apple"])
	fmt.Println(sweetFruits["lime"])

	// In cases where it is not sure if the key-value exists. if the check is not implemented then Go will return a
	// default value for the type e.g "" for strings, 0 for all numerics, false for bool.
	fruit, isPresent := sweetFruits["mango"]
	if isPresent {
		fmt.Println(fruit)
	} else {
		fmt.Println("Key not found")
	}
	fmt.Println()

	// Using for to print out the data
	// only keys
	for key := range sweetFruits {
		fmt.Println(key)
	}
	fmt.Println()

	// only values
	for _, value := range sweetFruits {
		fmt.Println(value)
	}
	fmt.Println()

	// both keys and values
	userAge := map[string]uint8{"Sam": 24, "Cornelius": 27, "Joseph": 24}
	for name, age := range userAge {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Change the value  using the key
	userAge["Sam"] = 250
	fmt.Println(userAge)

}
