/*
**************************************
Decrypt An Integer
Author: Camille Ollison
Date Completed: 12/3/25
Description: This program decrypts a character previously encrypted using a
siplified affine cipher algorithm
**************************************
*/
package main

import (
	"flag"
	"fmt"
)

func Decrypt(encrypted int, mult int, add int) string {
	var decryptedNum = (encrypted - add) / mult
	var decryptedChar string = string(decryptedNum)
	return decryptedChar
}
func main() {
	var encryptedInt int
	var multiplier int
	var adder int
	flag.IntVar(&encryptedInt, "encryptedInt", 0, "Specify an int")
	flag.IntVar(&multiplier, "multiplier", 0, "Specify an int")
	flag.IntVar(&adder, "adder", 0, "Specify and int")
	// Assign command line argument values to mapped variables
	flag.Parse()
	var decrypted string = Decrypt(encryptedInt, multiplier, adder)
	fmt.Println("Decrypted value:", decrypted)
}
