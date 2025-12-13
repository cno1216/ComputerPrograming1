package main

import "fmt"

func main() {

	var myInt int = 1

	// // if/else syntax
	// if myInt == 1 {
	// 	fmt.Println("One")
	// } else if myInt == 2 {
	// 	fmt.Println("Two")
	// } else if myInt == 3 {
	// 	fmt.Println("Three")
	// } else {
	// 	fmt.Println("Less than one or greater than three")
	// }

	// Switch statement syntax
	switch myInt {
		case 1:
			fmt.Println("One")
			fallthrough
		case 2:
			fmt.Println("Two")
			fallthrough
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Less than one or greater than three")
	}

}