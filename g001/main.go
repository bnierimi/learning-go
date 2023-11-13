package main

import "fmt"

func main(){
	// Integers: default value -> 0
	var intNum int = 20
	fmt.Println(intNum)

	// Floats: default value -> 0
	var floatNum = 3.3
	fmt.Println(floatNum)
	
	// Can't add the two types together so we convert one of them
	var addition = floatNum + float64(intNum)
	fmt.Println(addition)
	
	var division = floatNum / float64(intNum)
	fmt.Println(division)
	
	var multiplication = int(floatNum) % intNum
	fmt.Println(multiplication)

	// Strings: default value -> ""
	var my_name = "Radiance"
	fmt.Println(my_name)
	
	var my_story = `
	Well you know how it goes.....`
	fmt.Println(my_story)

	// Booleans: default value -> false
	var is_alive = true
	fmt.Println(is_alive)

	// Works pretty fine
	my_age := 12
	fmt.Println(my_age)

	var num1, num2 int = 1, 2
	// num1, num2 := 1, 2      // also
	fmt.Println(num1, num2)

	// Constants: immutable variable
	const MAX_MEMORY = 31.8
	fmt.Println(MAX_MEMORY)

	const AUTHOR = "Radiance Babajide"
	fmt.Println(AUTHOR)

	const pi = 3.1415
}