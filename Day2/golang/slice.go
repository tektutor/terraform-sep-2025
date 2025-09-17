package main

import "fmt"

func main() {
        //index             0  1  2  3  4  5
	intArray := [6]int{10,20,30,40,50,60}

	fmt.Println ( "Array elements are ...")
	fmt.Println(intArray)

	var mySlice []int = intArray[2:5] //2 is the lower bound index, while 5 is hte upper bound index, index 5 is not included

	fmt.Println("Slice elements are ...")
	fmt.Println(mySlice)

	//Let's modify the slice at certain indices
	//when the slice is modified, it will affect the original array that is pointed by the slice
	mySlice[0] = 100 //mySlice[0] is nothing but intArray[2]
	mySlice[1] = 200 //mySlice[1] is nothing but intArray[3]
	mySlice[2] = 300 //mySlice[2] is nothing but intArray[4]

	fmt.Println("Slice elements after modifying slice are ...")
	fmt.Println(mySlice)

	fmt.Println("Array elements after modifying slice are ...")
	fmt.Println(intArray)

	mySlice = append(mySlice, 400)
	fmt.Println("Array elements after appending 400 in slice are ...")
	fmt.Println(intArray)

	mySlice = append(mySlice, 500) //after this append, mySlice is no more associated/pointing to intArray
	fmt.Println("Array elements after appending 500 in slice are ...")
	fmt.Println(intArray)

	fmt.Println("Slice elements after appending 500 into slice are ...")
	fmt.Println(mySlice)

}
