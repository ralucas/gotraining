// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rj0QfwVPhX

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type Player struct {
	name   string
	atBats int64
	hits   int64
}

// Declare a method that calculates the batting average for a batter.
func (p Player) average() int64 {
	return p.hits / p.atBats
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.

	// Display the batting average for each player in the slice.
}
