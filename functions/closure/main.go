package main

import "fmt"

/* closure is a special type of anonymous function that references and retains access to variables from its surrounding scope,
even after the outer function has finished executing */

// counter() returns a function, not an int
func counter() func() int {

	// count is local to counter(). Normally, this would be destroyed when counter() finishes
	count := 0

	// returns an inner anonymous function (closure). This function "closes over" the variable `count`
	return func() int {

		// Each time this inner function is called, it modifies the SAME `count` variable in the outer scope
		count++

		// Return the updated value
		return count
	}
}

// More examples, taken from the Pro Go book
type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

var prizeGiveaway = false

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if prizeGiveaway {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}
		return price

	}
}

func main() {

	// counter() is executed ONCE here
	// It returns the inner function and assigns it to `result`
	// IMPORTANT: `count` is NOT destroyed — it is kept alive
	result := counter()

	// Each call invokes the SAME inner function
	// which still has access to the same `count`
	fmt.Println(result()) // 1
	fmt.Println(result()) // 2
	fmt.Println(result()) // 3
	fmt.Println(result()) // 4

	
	waterSportProducts := map[string]float64{
		"Kayak":       275,
		"Life jacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}
	prizeGiveaway = false
	waterCalc := priceCalcFactory(100, 0.2)
	prizeGiveaway = true
	soccerCalc := priceCalcFactory(50, 0.1)
	for product, price := range waterSportProducts {
		printPrice(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}

	calc := func(price float64) float64 {
		if price > 100 {
			return price + (price * 0.2)
		}
		return price
	}
	for product, price := range waterSportProducts {
		printPrice(product, price, calc)
	}
	calc = func(price float64) float64 {
		if price > 50 {
			return price + (price * 0.1)
		}
		return price
	}
	for product, price := range soccerProducts {
		printPrice(product, price, calc)
	}
}
