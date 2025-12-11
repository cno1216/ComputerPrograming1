/****************************************************************
arrays
Author: Matt
Date Completed: 11/09/2023
Description:
****************************************************************/

package main

import "fmt"

func main() {

	var x int = 42

	fmt.Println(x)

	// Properties of arrays
	// 1. Fixed size: you must know the size of the array when declaring 
	// 2. Data is stored contiguously (adjacent in memory)
	// 3. Homogeneous data (same type)
	// 4. Zero-based index (gotcha)

	var myArray [4]int

	fmt.Println(myArray)

	myArray[0] = 2
	myArray[1] = 4
	myArray[2] = 6
	myArray[3] = 8
	//myArray[4] = 666  - Compilation error - out of bounds

	fmt.Println(myArray[1])

	


}