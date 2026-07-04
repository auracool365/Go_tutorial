package main

import "fmt"

// The use of pointers(pass by reference) in struct methods allows us to modify the original struct data instead of working with a copy of it. This is
// particularly useful when we want to update the state of an object or when we want to avoid unnecessarily creating copies of large structs.

type User struct {
	name  string
	email string
	age   uint8
}

// Method with pointer receiver to update the user's email
func (user *User) updateEmail(newEmail string) {
	user.email = newEmail
}

// Method with pointer receiver to update the user's age
func (user *User) updateAge(newAge uint8) {
	user.age = newAge
}

// Method with pointer to modify all the fields of the struct, and return a pointer to the updated struct, allowing for method chaining.
func (user *User) updateUserInfo(newName, newEmail string, newAge uint8) *User {
	user.name = newName
	user.email = newEmail
	// only change age if the new age is one year older than the current age by 1, just to mimic age validation logic
	if newAge > user.age && newAge-user.age == 1 {
		user.age = newAge
	}
	return user
}

// without pointer receiver(pass by value), the method would receive a copy of the User struct.
func (user User) updateEmailWithoutPointer(newEmail string) {
	user.email = newEmail // vscode warns me of the duplication, since it does not update the original struct
}

// However, care must be taken when copying structs that contain pointer fields. The struct itself is copied, but the specific pointer field
// still references the same memory location. This means that changes to the data through one struct will affect the other struct that
// shares the same pointer.
type Profile struct {
	name string
	age  *int
}

func pointerFieldExample() {
	age := 25

	profile1 := Profile{
		name: "Alice",
		age:  &age,
	}

	fmt.Printf("My name is %s and I am %d years old.\n", profile1.name, *profile1.age) // Alice 25

	// Modifying the age through profile1's pointer
	*profile1.age = 26

	fmt.Printf("My name is %s and I am %d years old.\n", profile1.name, *profile1.age) // Alice 26

	// Creates a copy of profile1
	profile2 := profile1

	// Modify the age pointer type through profile2's pointer, unintentionally affecting profile1 as well
	*profile2.age = 30

	// As a value type, name is copied, so modifying profile2's name does not affect profile1's name
	profile2.name = "John"

	fmt.Printf("My name is %s and I am %d years old.\n", profile1.name, *profile1.age) // Alice 30
	fmt.Printf("My name is %s and I am %d years old.\n", profile2.name, *profile2.age) // John 30
}

// More complex struct with multiple pointer fields, slices, and maps
type Developer struct {
	name string
	age  int
}

type Account struct {
	score    *int
	nickname *string
	owner    *Developer
	tags     []string
	settings map[string]string
}

func main() {
	user := User{name: "Alice", email: "alice@example.com", age: 25}

	// pass by reference example
	fmt.Println("Before update:", user)
	user.updateEmail("alice.new@example.com")
	fmt.Println("After update:", user)

	fmt.Println("\nBefore age update:", user)
	user.updateAge(30)
	fmt.Println("After age update:", user)

	fmt.Println("\nBefore user info update:", user)
	user.updateUserInfo("Alice Smith", "alice.smith@example.com", 31)
	fmt.Println("After user info update:", user)

	// pass by value example
	fmt.Println("\nBefore update without pointer:", user)
	user.updateEmailWithoutPointer("alice.old@example.com")
	fmt.Println("After update without pointer:", user)

	// pointer field examples
	fmt.Println("\nPointer field example:")

	pointerFieldExample()

	userAge := 100
	profile3 := Profile{
		name: "John",
		age:  &userAge,
	}

	fmt.Printf("\nMy name is %s and I am %d years old.\n", profile3.name, *profile3.age)

	// Modifying the age through profile3's pointer
	*profile3.age = 101

	fmt.Printf("My name is %s and I am %d years old.\n", profile3.name, *profile3.age)

	profile4 := profile3

	// Modifying the age through profile4's pointer
	*profile4.age = 102
	fmt.Printf("My name is %s and I am %d years old.\n", profile3.name, *profile3.age)
	fmt.Printf("My name is %s and I am %d years old.\n", profile4.name, *profile4.age)

	// More complex struct with multiple pointer fields, slices, and maps
	score := 100
	nickname := "Ace"
	owner := &Developer{name: "Bob", age: 30}
	tags := []string{"golang", "programming", "developer"}
	settings := map[string]string{"theme": "dark", "notifications": "enabled"}

	account := Account{
		score:    &score,
		nickname: &nickname,
		owner:    owner,
		tags:     tags,
		settings: settings,
	}
	fmt.Printf("\nAccount details:\nScore: %d\nNickname: %s\nOwner: %s\nTags: %v\nSettings: %v\n",
		*account.score,
		*account.nickname,
		account.owner.name,
		account.tags,
		account.settings,
	)

	// Modifying account details through pointers
	*account.score = 150
	*account.nickname = "SuperAce"
	account.owner.age = 31
	account.tags = append(account.tags, "golang")
	account.settings["theme"] = "light"

	fmt.Printf("\nUpdated account details:\nScore: %d\nNickname: %s\nOwner: %s\nTags: %v\nSettings: %v\n",
		*account.score,
		*account.nickname,
		account.owner.name,
		account.tags,
		account.settings,
	)
}
