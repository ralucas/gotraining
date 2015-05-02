// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/-JBSUoux-v

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare and make a map of integer type values.
	var valMap = make(map[string]int)

	// Intialize some data into the map.
	valMap["age"] = 30
	valMap["days"] = 5
	valMap["num"] = 10
	valMap["chickens"] = 15
	valMap["money"] = 1000

	// Display each key/value pair.
	for i, name := range valMap {
		fmt.Println(i, ":", name)
	}
}
