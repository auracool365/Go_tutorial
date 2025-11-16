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
	if age >= 18 {
		return person.firstName + " can vote"
	}
	return person.firstName + " cannot vote"
}

func main() {
	person1 := Person{"John", "Doe", 100, true, CareerPath{"used in software examples"}, 1, Religion{"Christian"}}
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

	person3 := Person{"Sam", "Ade", 15, true, CareerPath{"Chef"}, 2, Religion{"Atheist"}}
	fmt.Println(canVote(person3, person3.age))
	fmt.Println(canVote(person1, person1.age))
	// fmt.Println(canVote(person_2, person_2.age)) // Error! since person_2 is not a Person struct. but the value person_2.age can be used

	var person4 Person
	person4.firstName = "Johnny"
	person4.lastName = "Joey"
	person4.age = 20
	person4.isMale = true
	person4.career = CareerPath{"Software Engineer"}
	person4.int = 3
	person4.religion = "Scientology"
	fmt.Println(person4)
	fmt.Println(person4.career.personCareer)

	fmt.Println(person4.getAge())
	person4.greet()

}
