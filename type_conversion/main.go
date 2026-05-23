package main

import (
	"fmt"
	"strconv"
)

// Advanced type conversion
type Car struct {
	isElectric bool
	wheels     int16
	price      float32
}

type Bike struct {
	isElectric bool
	wheels     int16
	price      float32
}

// Implicit conversion
func conversion() {
	// anonymous
	vehicle := struct {
		isElectric bool
		wheels     int16
		price      float32
	}{
		isElectric: true,
		wheels:     4,
		price:      35000.00,
	}

	var myCar Car
	var myBike Bike

	// myCar = myBike // Error: cannot use myBike (type Bike) as the type Car

	/* There is error when assigning Bike to Car, but not when assigning an anonymous type to Car. This is because Go's type system 
	treats named types and anonymous (literal) types differently. Named types require explicit conversion to ensure type safety, while 
	anonymous types can be assigned directly as long as their structure matches. */

	myCar = vehicle 

	fmt.Println("Car:", myCar)
	fmt.Println("Bike:", myBike)
	fmt.Println("Anonymous:", vehicle)
	fmt.Println()
}

// Explicit conversion
func explicitConversion() {
	var myCar Car
	var myBike Bike

	myCar = Car(myBike)

	fmt.Println("Car:", myCar)
	fmt.Println("Bike:", myBike)
	fmt.Println()
}

func signedUnsignedConversion() {
	var signedInt int
	var unsignedInt uint

	//signedInt = unsignedInt // Error: cannot use unsignedInt (type uint) as the type int
	signedInt = int(unsignedInt)

	fmt.Println("SignedInt:", signedInt)
	fmt.Println("UnsignedInt:", unsignedInt)
	fmt.Println()
}

func main() {
	// Casting an integer to a float
	intValue := 10
	floatValue := float64(intValue)
	fmt.Printf("Integer: %d, Float: %f\n", intValue, floatValue)

	// Casting a float to an integer
	floatValue2 := 3.14159
	intValue2 := int(floatValue2) // Casting float to int (truncating decimal part)
	fmt.Printf("Float: %f, Integer: %d\n", floatValue2, intValue2)

	// Casting a rune (int32) to a string
	runeValue := 'A'
	stringValue := string(runeValue)
	fmt.Printf("Rune: %c, String: %s\n", runeValue, stringValue)

	// Casting string to other types
	var myString = "10"
	var x = 10
	var y float32 = 2.0

	fmt.Println(myString, x, y)

	number, _ := strconv.Atoi(myString) // string to int
	fmt.Println(number)

	var result = number + 20
	fmt.Println(result)

	// integer to string
	fmt.Println("Sum: " + strconv.Itoa(number))

	// String to integer conversion
	str := "42"
	intValue, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String:", str, "Integer:", intValue)
	}

	// Integer to string conversion
	intValue2 = 123
	str2 := strconv.Itoa(intValue2)
	fmt.Println("Integer:", intValue2, "String:", str2)

	// Parsing a floating-point number from a string
	floatStr := "3.14159"
	floatValue, err = strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String:", floatStr, "Float:", floatValue)
	}

	// Other basic conversion examples
	var i int = 100
	var f float32 = float32(i)
	var u uint = uint(i)
	var s string = strconv.Itoa(i)
	fmt.Printf("Value: %v, type: %T\n", i, i)
	fmt.Printf("Value: %v, type: %T\n", f, f)
	fmt.Printf("Value: %v, type: %T\n", u, u)
	fmt.Printf("Value: %v, type: %T\n", s, s)
	fmt.Println()

	/* var a int = 42
	fmt.Printf("Value: %v, type: %T\n", a, a)

	var b float64 = a // error
	fmt.Printf("value: %v, type: %T\n", b, b)  */

	// Advanced 
	conversion()
	explicitConversion()
	signedUnsignedConversion()

	var pi float32 = 3.14159
	fmt.Println(pi)
	fmt.Printf("pi is of type %T\n", pi)
}
