// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/kzZZ24O23g

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import (
	"fmt"
)

// Declare a constant of kind string and assign a value.
const hello string = "Hola"

// Declare a constant of type integer and assign a value.
const num int = 10

// main is the entry point for the application.
func main() {
	// Display the value of both constants.
	fmt.Printf("Constants: %v %v\n", hello, num)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	const ki = 5
	const kfp = 2.12345

	var res = ki / kfp

	// Display the value of the variable.
	fmt.Println("Result:", res)
}
