// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/HPJvLJoupp

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import (
	"fmt"
)

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a composite literal.
	newUser := user{
		name:  "Jeff",
		email: "jeff@test.com",
		age:   45,
	}

	// Display the field values.
	fmt.Printf("NewUser: %T %#v\n", newUser, newUser)

	// Declare a variable using an anonymous struct.
	anon := struct {
		name  string
		email string
		age   int
	}{
		name:  "Chuck",
		email: "dude@dude.com",
		age:   25,
	}

	// Display the field values.
	fmt.Printf("Anon: %T %#v\n", anon, anon)
	fmt.Println("anon name", anon.name)
	fmt.Println("anon email", anon.email)
	fmt.Println("anon age", anon.age)
}
