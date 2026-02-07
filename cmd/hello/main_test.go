package main

import "fmt"

func ExampleMain() {

	main()

	// Output:
	// Hello Robbie.
	// Hello Frankie.
}

func ExampleGreet() {
	greet("Robbie")

	// Output:
	// Hello Robbie.
}

func ExampleDouble() {

	fmt.Println(double(2))
	fmt.Println(double(9))

	// Output:
	// 4
	// 18
}
