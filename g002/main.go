package main

import (
	"fmt"
	"errors"
)

func main(){
	name := "Mia"
	print(name)

	result, remainder, err := div(4, 2)

	// // Using if statements to handle errors
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("Result: %v", result)
	// } else {
		// 	fmt.Printf("Result: %v and Remainder: %v", result, remainder)
		// }
		
	// // Using switch-case statements to handle errors
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Result: %v", result)
	default:
		fmt.Printf("Result: %v and Remainder: %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was smooth")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
		
	}
}

func print(value string){
	fmt.Println(value)
}

func div(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("(x) Error: Cannot divide by Zero. Zero can't be the denominator")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}