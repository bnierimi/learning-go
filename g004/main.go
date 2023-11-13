package main

import (
	"fmt"
	"strings"
)

func main(){
	// Runes
	var char rune = 'a'
	fmt.Println(char)

	// Strings
	stringSlice := []string{"G", "o", "o", "d", "n", "e", "w", "s"}

	var stringBuilder strings.Builder
	for s := range stringSlice {
		stringBuilder.WriteString(stringSlice[s])
	}

	catString := stringBuilder.String()
	fmt.Printf("\n%v", catString)
}