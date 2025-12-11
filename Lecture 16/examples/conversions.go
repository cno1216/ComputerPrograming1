/****************************************************************
converstions
Author: Matt
Date Completed: 11/09/2023
Description:
****************************************************************/

package main

import (
	"fmt"
	"strconv"
)

// func FuncName(param1 int, param2 string) int, error {
// 	// body
// }



func main() {

	// Casting:

	// var myFloat float64 = 3.14
	// var myInt int

	// myInt = int(myFloat)

	// fmt.Println(myFloat)
	// fmt.Println(myInt)



	// Converting:


	var stringContainingNumber string = "42"

	var convertedInt int
	var err error

	convertedInt, err = strconv.Atoi(stringContainingNumber)

	// if no error occured
	if err == nil {
		fmt.Println(convertedInt)
	} else {
		fmt.Println(err)
	}

	
}