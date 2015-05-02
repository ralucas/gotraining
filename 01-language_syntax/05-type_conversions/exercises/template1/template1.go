// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Eg9m_rYm4V

// Declare a named type called counter with a base type of int. Declare and initalize
// a variable of this named type to its zero value. Display the value of this variable
// and the variables type.
//
// Declare a new variable of the named type assign it the value of 10. Display the value.
//
// Declare a variable of the same base type as your named typed. Attempt to assign the
// value of your named type variable to your new base type variable. Does the compiler
// allow the assignment?
package main

// Add imports.
import "fmt"

// Declare the counter named type with a base type of int.
type Counter int

func main() {
	// Declare and display the variable of the named type
	// to its zero value.
	var num Counter

	fmt.Printf("Number %T %v\n", num, num)

	// Declare and display the variable of the named type
	// to the value of 10.
	var ni Counter = 10

	fmt.Printf("Number %T %v\n", ni, ni)

	// Declare a variable of type int initialized to 1.
	var ti int = 1

	// Assign this new variable to one of the named variables.
	num = ti

	// Did you get a compiler error on the above assignment?
}
