package main

import ("fmt")

func main(){

	// Arrays
	var intArray [3]int
	fmt.Println(intArray)
	fmt.Println(intArray[0])
	intArray[1] = 123
	fmt.Println(intArray[1:3])

	// Memory location of an item in an Array
	fmt.Println(&intArray[0])
	
	// Initialising Arrays with values
	// var intArr2 [3]int = [3]int{1, 2, 3}
	intArray2 := [...]int{1, 2, 3}
	fmt.Println(intArray2)

	// Slices: are Arrays with no fixed length
	var intSlice []int = []int{0, 9, 8}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	intSlice2 := []int{6, 5}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Can also use make()
	var intSlice3 []int = make([]int, 3, 8)
	fmt.Println(intSlice3)

	// Maps: a set of Key:Value pairs
	dict := make(map[string]uint8)
	fmt.Println(dict)

	var dict2 map[string]uint8 = map[string]uint8{"Radiance":20, "Goodness":17}
	fmt.Println(dict2)

	var age, isInDict = dict["Holy"]
	// fmt.Printf("Age: %v | Is in dict: %v", age, isInDict)
	delete(dict, "Radiance")
	if isInDict {
		fmt.Printf("Age: %v", age)
	} else {
		fmt.Printf("(x) Invalid name : Name does not exist in dict\n")
	}

	// Loops in Go
	for person, age := range dict2{
		fmt.Printf("Name: %v\nAge: %v\n-----------\n", person, age)
	}

	// For looping through Arrays
	for index, value := range intArray2{
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	// Writing while loops
	i := 0
	for i < 10{
		fmt.Println(i)
		i++
	}
	
	// Other syntaxes
	for i:=0; i < 10; i++{
		fmt.Println(i)
	}
	for{
		if i >= 10{
			break
		}
		fmt.Println(i)
		i++
	}
}