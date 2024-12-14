package main

import (
	"fmt"
)

// "ch1/goTypes"

func main() {
	var name = "Felipe"
	var secondName = "Julia"

	name = "Kalieu"
	secondName = secondName + " Zen"

	const firstInitial rune = 'F'
	const lastInitial float32 = 'O'

	fmt.Println(secondName, firstInitial, lastInitial)

	if (name > secondName) {
		fmt.Println("Name is greater than second name")
	} else {
		fmt.Println("Second name is greater than name")
	}

}
