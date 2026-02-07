package main

import (
	"fmt"
)

/*
func double(x int) int {
	var doubled int
	doubled = x * 2
	return doubled
}
*/

func double(x int) int {
	doubled := x * 2
	return doubled
}

func greet(name string) {
	fmt.Printf("Hello %v.\n", name)
}

func main() {
	greet("Robbie")
	greet("Frankie")
}
