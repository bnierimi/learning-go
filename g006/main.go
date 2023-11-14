package main

import (
	"fmt"
)

func main(){
	// // Pointers
	// var p *int32 = new(int32)
	// var i int32
	
	// fmt.Printf("p points to: %v\n", *p)
	// fmt.Printf("And i is: %v", i)
	
	var p *int32
	var i int32

	i = 32
	p = &i // Get the memory address of a variable
	fmt.Printf("p is: %v and i is: %v", p, i)


	// Pointers and functions
	var thingOne = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nMemory Location of thingOne is: %p", &thingOne)
	var result [5]float64 = square(&thingOne)
	fmt.Printf("\nResult is: %v", result)
	fmt.Printf("\nThe value of thingOne is: %v", thingOne)
}

func square(thingTwo *[5]float64) [5]float64 {
	fmt.Printf("\nMemory Location of thingTwo is: %p", thingTwo)
	for i := range thingTwo {
		thingTwo[i] = thingTwo[i] * thingTwo[i]
	}
	return *thingTwo
}