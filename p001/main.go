package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "log"
)

// func input() string {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	var err = scanner.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf(">> %v", scanner.Text())
// 	return scanner.Text()
// }

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Welcome! %v\n", name)
	
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Get excited!")
	} else {
		fmt.Println("(x) Oops! sorry you are not old enough :(")
		return
	}

	fmt.Printf("Which is better, [O]ne Piece or [T]okyo Ghoul? ")
	var answer string
	fmt.Scan(&answer)

	if answer == "O" {
		fmt.Println(":) Correct!")
	} else if answer == "o" {
		fmt.Println(":) Correct!")
	} else {
		fmt.Println(":( Wrong!")
	}
}