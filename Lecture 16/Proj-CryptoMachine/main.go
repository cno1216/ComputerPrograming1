/***************************************
Crypto Machine
Author: Camille Ollison
Date Completed: 12/
Description: This program will read from a file and either encrypt or decrypt it.
***************************************/

package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
//    "flag"
// )

// objects
type Rotor struct {
	Multiplier int
	Adder      int
}

//Rotor member funcs

// pseudo
func NewRotor(m int, a int) Rotor {
	var r Rotor
	r.Multiplier = m
	r.Adder = a
	return r
}

func (r Rotor) GetMult() int {
	return r.Multiplier
}

func (r Rotor) GetAdder() int {
	return r.Adder
}

func (r Rotor) Encrypt(c string) int { //takes character and encrypts into int with mult and adder
	var encodedNum int = int(c[0])                   //cast into ascii
	return (encodedNum * r.GetMult()) + r.GetAdder() //return evaluated number
}

func (r Rotor) Decrypt(i int) string { //takes int and decrypts into character
	var decryptedNum = (i - r.GetAdder()) / r.GetMult()
	return string(decryptedNum) //return ascii cast into string
}

//obj

type CryptoMachine struct {
	Rotors []Rotor
}

//CryptoMachine member funcs

// pseudo
func NewCryptoMachine(r []Rotor) CryptoMachine {
	var cm CryptoMachine
	cm.Rotors = r
	return cm
}

func main() {

}
