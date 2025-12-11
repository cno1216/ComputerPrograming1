/****************************************************************
wordsInString
Author: Matt
Date Completed: 11/30/2023
Description:
****************************************************************/

package main

import "fmt"

func main() {

	var words []string
	var currentWord string

	var line string = "hello world and other words"

	for i := 0; i < len(line); i++ {

		if string(line[i]) == "o" {
			words = append(words, currentWord)
			currentWord = ""
		} else {
			currentWord += string(line[i])
		}
	}

	words = append(words, currentWord)

	fmt.Print(words)

}