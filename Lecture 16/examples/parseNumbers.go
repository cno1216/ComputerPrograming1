package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TwoInts struct {
	Int1 int
	Int2 int
} 

func NewTwoInts(int1 int, int2 int) TwoInts {
	var ti TwoInts
	ti.Int1 = int1
	ti.Int2 = int2
	return ti
}


func main() {

	// Open existing "FileToRead.txt" file in current directory and instantiate "File" object (readFile)
	readFile, err := os.Open("numbers.txt")

	// Check to see if any errors occurred while opening file.
	// If so, print error and bail out of function
	if err != nil {

		fmt.Println(err)
		return
	}

	// Instantiate new "Scanner" object from readFile object
	scanner := bufio.NewScanner(readFile)

	// Declare variable to hold data scanned from file
	var line string

	var allTheInts []TwoInts

	// Scan every line in the file
	for scanner.Scan() {

		// if successfully scanned, retrieve scanned text and store in variable
		line = scanner.Text()


		strings.Split(line, " ")

		var numbers []string = strings.Split(line, "`")

		firstInt, _ := strconv.Atoi(numbers[0])
		secondInt, _ := strconv.Atoi(numbers[1])

		//var parsedTwoInts TwoInts = 
		
		allTheInts = append(allTheInts, NewTwoInts(firstInt, secondInt))


		// // Print each word
		// for i := 0; i < len(words); i++ {
		// 	fmt.Println(words[i])
		// }
	}


	// Print each word
	for i := 0; i < len(allTheInts); i++ {
		fmt.Println(allTheInts[i])
	}

	
	// Close the file
	readFile.Close()
	
}