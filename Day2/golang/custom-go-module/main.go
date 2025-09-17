package main

import (
	"fmt"
	"addition/v2"
	"subtraction"
)

func main() {
	//x := 100.123 By default, golang will assume 100.123 as float64

	x := float32(100.123) // We are casting/converting float64 into float32
	y := float32(200.456) // We are casting/converting float64 into float32

	fmt.Println( "The sum of ", x, " and ", y, " is ", addition.Add( x, y ) )
	fmt.Println( "The difference of ", x , " and ", y, " is ", subtraction.Subtract( x, y ) )
}
