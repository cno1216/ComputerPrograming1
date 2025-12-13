/*
**************************************
Encrypt A Character
Author: Camille Ollison
Date Completed: 12/3/25
Description: This program encrypts a character using a siplified affine cipher
algorithm
**************************************
*/
package main

import (
	"flag"
	"fmt"
)

func Encrypt(c string, mult int, add int) int {
	var encodedNum int = int(c[0])
	var encryptedNum = (encodedNum * mult) + add
	return encryptedNum
}

func main() {
	var character string
	var multiplier int
	var adder int
	flag.StringVar(&character, "character", "A", "Specify one character")
	flag.IntVar(&multiplier, "multiplier", 1, "Specify an int")
	flag.IntVar(&adder, "adder", 1, "Specify and int")
	// Assign command line argument values to mapped variables
	flag.Parse()
	//
	var encrypted int = Encrypt(character, multiplier, adder)
	fmt.Println("Encrypted value:", encrypted)
}
