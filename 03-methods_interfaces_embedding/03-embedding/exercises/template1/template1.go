// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/a-Nzng_E6Z

// Declare a struct type named animal with two fields name and age. Declare a struct
// type named dog with the field bark. Embed the animal type into the dog type. Declare
// and initalize a value of type dog. Display the value of the variable.
//
// Declare a method named yelp to the animal type using a pointer reciever which displays the
// literal string "Not Implemented". Call the method from the value of type dog.
//
// Declare an interface named yelper with a single method called yelp. Declare a value of
// type yelper and assign the address of the value of type dog. Call the method yelp.
//
// Implement the yelper interface for the dog type. Be creative with the
// bark field. Call the method yelp again from the value of type yelper.
package main

// Add imports.
import (
	"fmt"
)

// Declare an interface type named yelper that has a
// single method called yelp().
type yelper interface {
	yelp()
}

// Declare a struct type named animal with two fields
// name of type string and age of type in.
type animal struct {
	name string
	age  int
}

// Declare a method for the animal struct that implements
// the yelper interface using a pointer receiver.
func (a *animal) yelp() {
	fmt.Printf("Yelping %v %v\n", a.name, a.age)
}

// Declare a struct type named dog that embeds the animal
// type and has a field named bark of type int.
type dog struct {
	animal
	bark int
}

// Declare a method for the dog struct that implements
// the yelper interface using a pointer receiver.
func (d *dog) yelp() {
	fmt.Printf("Yelping %v\n", d.bark)
}

// main is the entry point for the application.
func main() {
	// Declare and initialize a variable of type dog.
	spot := dog{
		animal: animal{
			name: "Spot",
			age:  14,
		},
		bark: 6,
	}

	// Display the value of the variable.
	fmt.Println(spot)

	// Declare a variable of the yelper interface type.
	var y yelper

	// Assign the dog variable to the iterface variable.
	// The call the interface method.
	y = spot

	y.yelp()
}
