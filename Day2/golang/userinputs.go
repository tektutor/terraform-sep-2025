package main

import "fmt"

func main() {
	//The below line declares variables named x and y, later it is initialized with value 0
	x := 0
	y := 0
	
	fmt.Print("Enter your first integer :" )
	fmt.Scanf("%d", &x )

	fmt.Print("Enter your second integer :" )
	fmt.Scanf("%d", &y )

	fmt.Println("Value of x :", x )
	fmt.Println("Value of y :", y )

	//Declares a variable named tmp of type string
	var temp string
	fmt.Println("Press any key to exit ...")
	fmt.Scanln(&temp)
}
