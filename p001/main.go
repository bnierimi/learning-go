package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func input() string {
	var reader = bufio.NewReader(os.Stdin)
	var line, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">> %v", line)
	return line
}

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
	var answer = input()
	// fmt.Scan(&answer)
	fmt.Println(answer)

	// if answer == "O" {
	// 	fmt.Println(":) Correct!")
	// } else if answer == "o" {
	// 	fmt.Println(":) Correct!")
	// } else {
	// 	fmt.Println(":( Wrong!")
	// }
}