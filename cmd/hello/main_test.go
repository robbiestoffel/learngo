package main

import "fmt"

func ExampleGreet() {
	fmt.Println(Greet("World"))
	fmt.Println(Greet("Frankie"))

	// Output:
	// Hello World.
	// Hello Frankie.
}

func ExampleDouble() {

	fmt.Println(Double(2))
	fmt.Println(Double(9))

	// Output:
	// 4
	// 18
}
