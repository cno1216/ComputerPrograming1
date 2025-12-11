/****************************************************************
betterSlices
Author: Matt
Date Completed: 11/09/2023
Description:
****************************************************************/

package main

import "fmt"

func main() {

	var mySlice []int
	fmt.Println(mySlice)


	mySlice = append(mySlice, 42)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 11)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 60)
	fmt.Println(mySlice)

	otherSlice := append(mySlice, 941)
	fmt.Println(mySlice)
	fmt.Println(otherSlice)

	
	// Print out all of the values in the slice
	for i := 0; i < len(otherSlice); i++ {

		fmt.Print("[", i, "] = ", otherSlice[i], "\n")
	}

}