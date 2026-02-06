package main

import (
	"fmt"
	"log"
)

// this isn't a function, but don't worry about it
func main() {
	// a block of code is a set of things that run at the same level, in
	// this case, a block of code is between the curly brackets
	println("debug: hello, friend.")
	println("debug: again")
	fmt.Println("Hello, 🌍")
	log.Println("hello with log.") // best way to print to standard error

	fmt.Printf("%v\n", "Hello again")
	fmt.Printf("%v\n", 123.23/234)
	// %v - verbose. converts any type into a string to print it.
	fmt.Printf("%U\n", '🌍') // running this as a number so that we can see what the unicode is
	fmt.Printf("%v\n", string('\U0001F30D'))

	fmt.Printf("%b\n", []byte("Hello again"))
}
