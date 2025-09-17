package main

import "fmt"

func main() {
	count := 5 //Declares a count variable of type int and assigns a value 5

	//similar to while loop
	for count > 0 {
		fmt.Println("Before decrementing ", count)

		count-- //equivalent to count = count - 1
		//golang doesn't support pre-decrement or pre-increment unlinke C and C++

		fmt.Println("After decrementing ", count )
	}
	fmt.Println("Value of count is ", count, " after for loop")

	count = 0 //variable is already declared in line number 7, we are just trying to reset the value to 0 here

	fmt.Println()
	//Regular for loop
	for count=1; count<10; count++ {
		fmt.Printf( "%d\t", count )
	}
	fmt.Println()

	//similar to while do
	count = 0

	for {  //Infinite loop

		fmt.Printf ("Inside for loop %d\n", count ) 	
		count++

		if count > 3 {
			break
		}
	}
	fmt.Println("Control reached outside for infinite for loop")
}
