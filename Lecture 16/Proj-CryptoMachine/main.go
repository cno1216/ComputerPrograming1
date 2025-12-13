/***************************************
Crypto Machine
Author: Camille Ollison
Date Completed: 12/
Description: This program will read from a file and either encrypt or decrypt it.
***************************************/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func (r Rotor) Cipher(c int) int { //takes character and encrypts into int with mult and adder
	return (c * r.GetMult()) + r.GetAdder() //return evaluated number
}

func (r Rotor) Decipher(i int) int { //takes int and decrypts into character
	return (i - r.GetAdder()) / r.GetMult()

}

//obj

type CryptoMachine struct {
	Rotors []Rotor
}

//CryptoMachine member funcs

// pseudo
func NewCryptoMachine() CryptoMachine {
	var cm CryptoMachine

	// Open existing "rotors.txt" file in current directory and instantiate "File" object (readFile)
	readFile, err := os.Open("rotors.txt")

	// Check to see if any errors occurred while opening file.
	// If so, print error and bail out of function
	if err != nil {
		fmt.Println(err)
		return cm
	}

	// Instantiate new "Scanner" object from readFile object
	scanner := bufio.NewScanner(readFile)

	// Declare variable to hold data scanned from file
	var line string
	var rotors []Rotor

	// Scan every line in the file
	for scanner.Scan() {
		// if successfully scanned, retrieve scanned text and store in variable
		line = scanner.Text()

		var numbers []string = strings.Split(line, " ")
		mult, _ := strconv.Atoi(numbers[0])
		add, _ := strconv.Atoi(numbers[1])

		rotors = append(rotors, NewRotor(mult, add))
	}

	cm.Rotors = rotors

	// Close the file
	readFile.Close()

	return cm
}

func (cm CryptoMachine) GetRotors() []Rotor {
	return cm.Rotors
}

//non member funcs

func RunEncryption(cm CryptoMachine, message string) {
	var i int
	for i = 0; i < len(message); i++ {
		var character int = int(message[i]) //cast character into ascii
		for j := 0; j < len(cm.GetRotors()); j++ {
			character = cm.GetRotors()[j].Cipher(character) //encrypt the character with j rotor and assign it
		}
		fmt.Print(character, " ")
	}

}

func RunDecryption(cm CryptoMachine, message string) {
	var characters []string = strings.Split(message, " ")

	for i := 0; i < len(characters); i++ {
		character, _ := strconv.Atoi(characters[i])
		for j := len(cm.GetRotors()) - 1; j >= 0; j-- {
			character = cm.GetRotors()[j].Decipher(character)
		}
		fmt.Print(string(character))
	}
}

func main() {
	var cryptograph = NewCryptoMachine()
	var mode string
	var message string

	flag.StringVar(&mode, "mode", "", "Specify a string for mode")
	flag.StringVar(&message, "message", "", "Specify a string for message")

	//assign flags
	flag.Parse()

	switch mode {
	case "encrypt":
		RunEncryption(cryptograph, message)
		fallthrough
	case "decrypt":
		RunDecryption(cryptograph, message)
	default:
		fmt.Print("Error")
	}

	// for i := 0; i < len(cryptograph.GetRotors()); i++ {
	// 	fmt.Println(cryptograph.GetRotors()[i])
	// }

	// fmt.Println(mode)
	// fmt.Println(message)

}
