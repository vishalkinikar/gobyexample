/*
	Go has various value types including
		strings,
		integers,
		floats,
		booleans, etc.

	Here are a few basic examples.
*/

package main

import (
	"fmt"
)

func main() {

	// Strings, which can be added together with +.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)

	// Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
