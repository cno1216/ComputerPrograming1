/***************************************
Extreme Points
Author: Camille Ollison
Date Completed: 12/10/25
Description: This program will list the coordinate points in a file and minimum and maximum X & Y points.
***************************************/
/*
object oriented.
GetPoints will -> parse and gen a Point obj from each line -> use slice to hold all ->v
-> close file -> return slice
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	Xcord int
	Ycord int
}

// Member func
func NewPoint(x int, y int) Point {
	var p Point
	p.Xcord = x
	p.Ycord = y
	return p
}

func (p Point) GetX() int {
	return p.Xcord
}

func (p Point) GetY() int {
	return p.Ycord
}

func (p Point) Print() {
	fmt.Print("(", p.Xcord, ", ", p.Ycord, ")")
}

// Non Member func
func FindMin(p []Point, coord string) Point {
	var minimum Point = p[0] //start at first value

	if coord == "x" || coord == "X" { //checking x values
		for i := 0; i < len(p); i++ { //if min is greater than the one being checked
			if minimum.GetX() > p[i].GetX() {
				minimum = p[i]
			} //set min to current
		}
	} else { //checking y values
		for i := 0; i < len(p); i++ { //if min is greater than the one being checked
			if minimum.GetY() > p[i].GetY() {
				minimum = p[i]
			} //set min to current
		}
	}

	return minimum
}

func FindMax(p []Point, coord string) Point {
	var maximum Point = p[0] //start at first value

	if coord == "x" || coord == "X" { //checking x values
		for i := 0; i < len(p); i++ { //if max is less than the one being checked
			if maximum.GetX() < p[i].GetX() {
				maximum = p[i]
			} //set max to current
		}
	} else { //checking y values
		for i := 0; i < len(p); i++ { //if max is less than the one being checked
			if maximum.GetY() < p[i].GetY() {
				maximum = p[i]
			} //set max to current
		}
	}

	return maximum
}

func GetPoints() []Point {

	// Declare variable to hold data scanned from file
	var line string
	var points []Point

	// Open existing "points.txt" file in current directory and instantiate "File" object (readFile)
	readFile, err := os.Open("points.txt")

	// Check to see if any errors occurred while opening file.
	// If so, print error and bail out of function
	if err != nil {
		fmt.Println(err)
		return points
	}

	// Instantiate new "Scanner" object from readFile object
	scanner := bufio.NewScanner(readFile)

	// Scan every line in the file
	for scanner.Scan() {

		// if successfully scanned, retrieve scanned text and store in variable
		line = scanner.Text()

		//split line into two strings
		var coordinates []string = strings.Split(line, ",")

		var x int
		var y int

		x, err = strconv.Atoi(coordinates[0])
		// if error occured
		if err != nil {
			fmt.Println(err)
			return points
		}
		y, err = strconv.Atoi(coordinates[1])
		// if error occured
		if err != nil {
			fmt.Println(err)
			return points
		}

		//add line to points
		points = append(points, NewPoint(x, y))

	}

	// Close the file
	readFile.Close()

	return points
}

func main() {

	var points []Point = GetPoints()
	fmt.Print(len(points), " coordinate points found:\n\n")
	for i := 0; i < len(points); i++ {
		points[i].Print()
		fmt.Print("\n")
	}

	var minX Point = FindMin(points, "x")
	var maxX Point = FindMax(points, "x")
	var minY Point = FindMin(points, "y")
	var maxY Point = FindMax(points, "y")

	fmt.Print("\nMinimum X Coordinate Point: ")
	minX.Print()
	fmt.Print("\nMaximum X Coordinate Point: ")
	maxX.Print()
	fmt.Print("\nMinimum Y Coordinate Point: ")
	minY.Print()
	fmt.Print("\nMaximum Y Coordinate Point: ")
	maxY.Print()

}
