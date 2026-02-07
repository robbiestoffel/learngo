package main

import (
	"fmt"
)

func main() {
	var name = "Robbie"
	fmt.Println("Hello " + name + ".")
	// -or-
	name = "Stoff"
	fmt.Printf("Hello %v.\n", name)
}
