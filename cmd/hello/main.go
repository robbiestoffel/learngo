package main

import "fmt"

/*
func double(x int) int {
	var doubled int
	doubled = x * 2
	return doubled
}
*/

/*
func greet(name string) {
	fmt.Printf("Hello %v.\n", name)
}
*/

// Math functions
func Double(x int) int           { return x * 2 }
func Half(x int) int             { return x / 2 }
func Increment(x int, y int) int { return x + y }
func Decrement(x int, y int) int { return x - y }

func Greet(name string) string { return "Hello " + name + "." }

func main() {
	fmt.Println(Greet("Robbie"))
	fmt.Println(Greet("Frankie"))
	fmt.Println(Double(3))
	fmt.Println("Half of 256 is", Half(256))

}
