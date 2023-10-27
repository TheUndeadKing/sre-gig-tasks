package main

import (
	"fmt"
	comparevalue "tasktwo/compareValue"
)

func main() {
	// Generate a random number between 1 and 100
	// A: Created custom func inside generatevalue sub dir

	// If the number is higher than 50 print "It's closer to 100"
	// If the number is lower than 50 print "It's closer to 0"
	// Print the generated random number

	i, j := comparevalue.CompareValue()

	fmt.Println(i)
	fmt.Println("Random Value: ", j)

}
