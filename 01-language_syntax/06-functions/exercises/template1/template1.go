// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/p0vlsW5sVL

// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type User struct {
	name string
	age  int
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.
func createUser() (*User, error) {
	// Create a value of type user and return the proper values.
	var u User

	return &u{"Richard", 37}, nil
}

// main is the entry point for the application.
func main() {
	// Use the function to create a value of type user. Check
	// the error being returned.
	u2, err := createUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value that the pointer points to.
	fmt.Println(u2)

	// Call the function again and just check the error.
}
