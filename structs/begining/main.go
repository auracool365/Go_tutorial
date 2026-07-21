package main

import "fmt"

// Struct allows bundling together data that relate to each other in usage, essentially creating a new data type
type Person struct {
	firstName string
	lastName  string
	age       uint8
	isMale    bool
	career    CareerPath

	// properties can be added just by type, without a name
	int      //id
	Religion //religion
}

type CareerPath struct {
	personCareer string
}

type Religion struct {
	religion string
}

// Methods(The syntax is different from normal functions because methods are bound to the receiver )
func (person Person) greet() {
	fmt.Println("Hello", person.firstName, person.lastName)
}

func (person Person) getAge() (string, uint8, string) {
	userName := fmt.Sprintf("%s %s is", person.firstName, person.lastName)
	return userName, person.age, "years old"
}

// This is a normal function and cannot be attached to the struct objects.
func canVote(person Person, age uint8) string {
	if age >= 18 && age <= 90 {
		return fmt.Sprintf("at age %d, %s can vote", age, person.firstName)
	} 
	if age > 90 {
		return fmt.Sprintf("at age %d, %s can vote, but may be too old to exercise that right", age, person.firstName)
	}
	return fmt.Sprintf("at age %d, %s cannot vote", age, person.firstName)
}

// Difference between methods and normal functions is that methods are bound to the receiver, which is the struct object, and can access 
// its fields and other methods. Normal functions are not bound to any receiver and cannot access the fields of a struct directly. Methods can be 
// called on struct instances, while normal functions are called independently.

type Rectangle struct {
	Width  float64
	Height float64
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method to calculate the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Normal function to calculate the area of the rectangle
func CalculateArea(r Rectangle) float64 {
	return r.Width * r.Height
}

// Normal function to calculate the perimeter of the rectangle
func CalculatePerimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Just a normal function 
func add(a, b int) int {
	return a + b
}

// Create a struct of products and their prices, Store adds items to database, user can add items to cart, and
// checkout to get the total price of the items in the cart, with the option to remove items from the cart before
// checkout. quantity, discounts, etc.

// Product struct
type Product struct {
	name  string
	price float64
}

// Cart struct to hold products added by the user
type Cart struct {
	products []Product
}

// store struct(database using a slice to hold products)
type Store struct {
	products []Product
}

// Add product to store
func (s *Store) addProduct(product Product) {
	s.products = append(s.products, product)
}

// Remove product from store
func (s *Store) removeProduct(productName string) {
	for i, product := range s.products {
		if product.name == productName {
			s.products = append(s.products[:i], s.products[i+1:]...)
			return
		}
	}
}

// Add product to cart
func (c *Cart) addToCart(product Product) {
	c.products = append(c.products, product)
}

// Remove product from cart
func (c *Cart) removeFromCart(productName string) {
	for i, product := range c.products {
		if product.name == productName {
			c.products = append(c.products[:i], c.products[i+1:]...)
			return
		}
	}
}

// Checkout and get total price of items in cart
func (c *Cart) checkout() float64 {
	var total float64
	for _, product := range c.products {
		total += product.price
	}
	return total
}

func main() {
	person1 := Person{
		"John",
		"Doe",
		100,
		true,
		CareerPath{"Software Engineer"},
		1,
		Religion{"Christian"},
	}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person1.age)

	person1.lastName = "Thompson" // modifiable from the outside
	fmt.Println(person1)

	// A struct can be created on the go for a one-time use, so no reusability
	person2 := struct {
		firstName string
		lastName  string
		age       uint8
		isMale    bool
		career    CareerPath
	}{firstName: "Jane", // Used immediately
		lastName: "Doe",
		age:      20,
		isMale:   false,
		career:   CareerPath{"Also used in software examples"},
	}
	fmt.Println(person2)
	fmt.Println(person2.firstName)

	person1.greet()
	fmt.Println(person1.getAge())

	if person2.isMale {
		fmt.Println(person2.firstName, "is a male")
	} else {
		fmt.Println(person2.firstName, "is a female")
	}

	person3 := Person{
		firstName: "Sam",
		lastName:  "Ade",
		age:       15,
		isMale:    true,
		career:    CareerPath{"Chef"},
		int:       2,
		Religion:  Religion{"Atheist"},
	}
	fmt.Println(canVote(person3, person3.age))
	fmt.Println(canVote(person1, person1.age))
	// fmt.Println(canVote(person_2, person_2.age)) // Error! since person_2 is not a Person struct. but the value person_2.age can be used

	// Creating a struct without assigning values to the fields, the fields will be initialized with the default value for their type
	var person4 Person
	fmt.Println(person4)
	person4.firstName = "Johnny"
	person4.lastName = "Joey"
	person4.age = 20
	person4.isMale = true
	person4.career = CareerPath{"Software Engineer"}
	person4.int = 3
	person4.religion = "Scientology" // The field have no name in the struct, it is assigned directly, specifying Religion() is error
	fmt.Println(person4)
	fmt.Println(person4.career.personCareer)

	fmt.Println(person4.getAge())
	person4.greet()

	// Product, Store and Cart
	// Field names can be ignored if the values are assigned in the same order as they are defined in the
	// struct (Not recommended due to poor readability and error-prone if mixed up)
	product1 := Product{"Laptop", 999.99}
	product2 := Product{"Smartphone", 499.99}
	product3 := Product{"Headphones", 199.99}
	product4 := Product{"Monitor", 299.99}
	product5 := Product{"Keyboard", 99.99}
	product6 := Product{"Mouse", 49.99}

	store := Store{}
	store.addProduct(product1)
	store.addProduct(product2)
	store.addProduct(product3)
	store.addProduct(product4)
	store.addProduct(product5)
	store.addProduct(product6)
	fmt.Println("\nStore products:", store.products)
	store.removeProduct("Headphones")
	fmt.Println("\nStore products after removal:", store.products)

	cart := Cart{}
	cart.addToCart(product1)
	cart.addToCart(product2)
	fmt.Println("\nCart products:", cart.products)

	total := cart.checkout()
	fmt.Printf("Total price: $%.2f\n", total)

	cart.removeFromCart("Laptop")
	fmt.Println("\nCart products after removal:", cart.products)

	total = cart.checkout()
	fmt.Printf("Total price after removal: $%.2f\n", total)

	// Demonstrating the use of methods and normal functions to calculate area and perimeter of a rectangle
	rect := Rectangle{Width: 5, Height: 10}

	// Using methods bound to the Rectangle struct
	fmt.Printf("\nUsing methods:\n")
	fmt.Printf("Area of rectangle: %.2f\n", rect.Area())
	fmt.Printf("Perimeter of rectangle: %.2f\n", rect.Perimeter())

	// Using normal functions on the struct, but not bound to the Rectangle struct
	fmt.Printf("\nUsing normal functions:\n")
	fmt.Printf("Area of rectangle: %.2f\n", CalculateArea(rect))
	fmt.Printf("Perimeter of rectangle: %.2f\n", CalculatePerimeter(rect))

	// normal add() function
	a, b := 5, 10
	sum := add(a, b)
	fmt.Printf("\nThe sum of %d and %d is: %d\n", a, b, sum)

}

