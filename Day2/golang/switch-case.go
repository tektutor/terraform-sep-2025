package main

import "fmt"

func main() {

	var direction string

	fmt.Println("Possible values are east,west,south,north")

	fmt.Print("Enter some direction :")
	fmt.Scanln(&direction)

	switch direction {
	case "east":
		fmt.Println("You enterred direction ", direction)
	case "west":
		fmt.Println("You enterred direction ", direction)
	case "south":
		fmt.Println("You enterred direction ", direction)
	case "north":
		fmt.Println("You enterred direction ", direction)
	default:
		fmt.Println("Invalid direction", "possible values are east, west, north, south")
	}
}
