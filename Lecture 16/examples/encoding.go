/****************************************************************
encoding
Author: Matt
Date Completed: 11/09/2023
Description:
****************************************************************/

package main

import "fmt"

func main() {

	var stringContainingCharacter string = "A Happy Day"

	var encodedNum int = int(stringContainingCharacter[2])

	fmt.Println(encodedNum)


	// var asciiInt int = 42
	// var encodedCharacter string = string(asciiInt)

	// fmt.Println(encodedCharacter)
	
}