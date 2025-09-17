package main

import (
   "fmt"
   "time"
)

func firstFunction( count int ) {
   for i := 0; i < count; i++ {
	fmt.Println("First function ", i )
	time.Sleep(time.Millisecond * 5 )
   }
}

func secondFunction( count int ) {
   for i := 0; i < count; i++ {
	fmt.Println("Second function ", i )
	time.Sleep(time.Millisecond * 5 )
   }
}

func main() {
	fmt.Println ( "Press any key to exit ...")

	//Invoking firstFunction and secondFunction in sequence one after the other
	firstFunction(10)
	secondFunction(10)

	//We wish to run both firstFunction and secondFunction in parallel
	go firstFunction( 1000 )
	go secondFunction( 1000 )

	var tmp string
	fmt.Scanln(&tmp) //this will make sure the program waits until some key is pressed to exit
}
