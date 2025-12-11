package main

import (
	"bufio"
	"fmt"
	"os"
)

// Write to file example
func WriteToFile() bool {

	// Create/truncate "data.txt" file in current directory and instantiate "File" object (writeFile)
	writeFile, fileCreateErr := os.Create("data.txt")

	// Check to see if any errors occurred while creating file.
	// If so, print error and bail out of function
	if fileCreateErr != nil {

		fmt.Println(fileCreateErr)
		return false
	}
	
	
	// Instantiate new "Writer" object from writeFile object
	writer := bufio.NewWriter(writeFile)

	// Write some text to the writer object
	_, fileWriteErr := writer.WriteString("I'm writing this to the file")

	// Check to see if any errors occurred while writing the string
	// to the writer object (i.e. write buffer).
	// If so, print error and bail out of function
	if fileWriteErr != nil {
				
		fmt.Println(fileWriteErr)
		return false
	}

	// Flush the data in the writer object (i.e. write buffer) to the actual file
	writer.Flush()

	// Close the file
	writeFile.Close()

	return true
}


// Read from file example
func ReadFromFile() bool {

	// Open existing "data.txt" file in current directory and instantiate "File" object (readFile)
	readFile, err := os.Open("data.txt")

	// Check to see if any errors occurred while opening file.
	// If so, print error and bail out of function
	if err != nil {

		fmt.Println(err)
		return false
	}

	// Instantiate new "Scanner" object from readFile object
	scanner := bufio.NewScanner(readFile)

	// Declare variable to hold data scanned from file
	var dataFromFile string

	// Scan the first line from the file
	if scanner.Scan() {

		// if successfully scanned, retrieve scanned text and store in variable
		dataFromFile = scanner.Text()
	
	} else {

		fmt.Println("Unable to scan line from data.txt")
		return false
	}

	// Display scanned text to terminal
	fmt.Println("Data from file:", dataFromFile)

	// Close the file
	readFile.Close()

	return true
}



func main() {

	if WriteToFile() {
		ReadFromFile()
	}
}