package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Open existing "FileToRead.txt" file in current directory and instantiate "File" object (readFile)
	readFile, err := os.Open("FileToRead.txt")

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

	// Scan every line in the file
	for scanner.Scan() {

		// if successfully scanned, retrieve scanned text and store in variable
		line = scanner.Text()


		strings.Split(line, " ")

		var words []string = strings.Split(line, " ")
		// var currentWord string

		// for i := 0; i < len(line); i++ {

		// 	if string(line[i]) == " " {
		// 		words = append(words, currentWord)
		// 		currentWord = ""
		// 	} else {
		// 		currentWord += string(line[i])
		// 	}
		// }

		// words = append(words, currentWord)


		// Print each word
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}

	
	// Close the file
	readFile.Close()
	
}