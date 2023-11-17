// func main() {
//     fmt.Println("input text:")
//     scanner := bufio.NewScanner(os.Stdin)
//     scanner.Scan()
//     err := scanner.Err()
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Printf("read line: %s-\n", scanner.Text())
// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func input(text *string) {
	// Read line of input: Getting user's input
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("You entered: %v\n", scanner.Text())
	*text = scanner.Text()
}

func main() {
	var response string
	fmt.Printf("Enter text: ")
	input(&response)
	fmt.Printf("Response: %v", response)
}
