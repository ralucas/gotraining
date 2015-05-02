// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/1xUWjHMB3I

// Declare three variables that are initalized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initalize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.

	// Display the value of those variables.

	// Declare variables and initalize.
	// Using the short variable declaration operator.

	// Display the value of those variables.

	// Perform a type conversion.

	// Display the value of that variable.

	var first int
	var second string
	var third bool

	fmt.Printf("first: %T %v\n", first, first)
	fmt.Printf("second: %T %v\n", second, second)
	fmt.Printf("third: %T %v\n", third, third)

	hi := "Hi"
	num := 50
	b := true

	fmt.Println(hi)
	fmt.Println(num)
	fmt.Println(b)

	convo := float32(3.14)

	fmt.Printf("convo: %T %v\n", convo, convo)

	convoTwo := "Hola"

	fmt.Printf("convoTwo: %T %v\n", convoTwo, convoTwo)
}
