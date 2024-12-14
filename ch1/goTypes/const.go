package goTypes

import "fmt"

// Constants in Go are a way to give names to literals. There is NO way in Go to declare that a variable is immutable

// Constants can be typed or untyped. An untyped constant works exactly like a literal;
// it has no type of its own but does have a default type that is used when no other type can be inferred.

func UsingConst() {
	a := 5
	b := 10

	// Go doesn't provide a way to specify that a value calculated at runtime is immutable.
	const sum = a + b // This won't compile! 

	fmt.Println(sum)
}