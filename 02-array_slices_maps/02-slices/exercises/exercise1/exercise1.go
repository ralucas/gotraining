// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/prAjmbCrVe

// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of the first and second index
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var numbers []int

	// Appens numbers to the slice.
	for i := 0; i < 5; i++ {
		numbers = append(numbers, i*10)
	}

	// Display each value.
	for _, number := range numbers {
		fmt.Println(number)
	}

	// Declare a slice of strings.
	names := []string{"Bill", "Lisa", "Jim", "Cathy"}

	// Display each index position and name.
	for i, name := range names {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}

	// Take a slice of the first and second index.
	slice := names[1:3]

	// Display the value of the new slice.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}
