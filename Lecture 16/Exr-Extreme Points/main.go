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

import "fmt"

type Point struct {
	Xcord int
	Ycord int
}

func NewPoint(x int, y int) Point{
	var p Point
	p.Xcord = x
	p.Ycord = y
	return p
}

func GetPoints() {

}


func main() {
   






}