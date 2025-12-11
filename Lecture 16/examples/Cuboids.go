package main

import "fmt"

// Custom data type definition
type Cuboid struct {
	Width int
	Length int
	Height int
}

// Value-reciever Member function that prints cuboid data
func (c Cuboid) Print(cuboidNum int) {
	fmt.Print("Cuboid ", cuboidNum, ":\n")
	fmt.Print("Width: ", c.Width, ", Length: ", c.Length, ", Height: ", c.Height, "\n\n")
}

// Pointer-reciever member function that doubles the cuboid dimensions
func (c *Cuboid) DoubleIt() {
	c.Width *= 2
	c.Length *= 2
	c.Height *= 2
}


// Non-member function that prints cuboid data
func PrintCuboid(c Cuboid, cuboidNum int) {

	fmt.Print("Cuboid ", cuboidNum, ":\n")
	fmt.Print("Width: ", c.Width, ", Length: ", c.Length, ", Height: ", c.Height, "\n\n")

	c.Width = 7
	fmt.Print("Cuboid ", cuboidNum, ":\n")
	fmt.Print("Width: ", c.Width, ", Length: ", c.Length, ", Height: ", c.Height, "\n\n")
}

// Calculates and returns Cuboid volume
func CuboidVolume(c Cuboid) int {
	return c.Width * c.Length * c.Height
}


// Pseudo-constructor
func NewCuboid(width int, length int, height int) Cuboid {
	var c Cuboid
	c.Width = width
	c.Length = length
	c.Height = height
	return c
}



func main() {

	// Instantiate and initialize first Cuboid object
	//var cuboid1 Cuboid = NewCuboid(3, 4, 2)
	
	// Instantiate and initialize second Cuboid object
	var cuboid2 Cuboid = NewCuboid(5, 1, 6)
	
	//PrintCuboid(cuboid1, 1) // calling non-member function

	
	//fmt.Print("Width: ", cuboid1.Width)

	cuboid2.Print(2) // calling member function

	cuboid2.DoubleIt()

	cuboid2.Print(2)

	//PrintCuboid(cuboid2, 2)
	

	// fmt.Println("Cuboid 1 Volume:", CuboidVolume(cuboid1)) //cuboid1.Width * cuboid1.Length * cuboid1.Height)
	// fmt.Println("Cuboid 2 Volume:", CuboidVolume(cuboid2)) //cuboid2.Width * cuboid2.Length * cuboid2.Height)
}
