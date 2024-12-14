package main

import "ch1/goTypes"

const myNum int64 = 10

// This is a constant block, defining multiple constants
const (
	idKey   = "id"
	nameKey = "name"
)

const sum = 20 * 10
const whatever string = "whatever"

func main() {
/*
	const phrase = "hello"

	fmt.Println(myNum)
	fmt.Println(phrase)



	//myNum = myNum + 1 // It'll not compile
	//phrase = "Hi, how are you going" 

	fmt.Println(myNum)
	fmt.Println(phrase)
	fmt.Println(whatever)
	fmt.Println(sum)
*/


	goTypes.UsingConst()


}



















