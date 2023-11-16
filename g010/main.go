package main

import (
	"fmt"
)

// func main() {
// 	var intSlice = []int{1, 2, 3}
// 	fmt.Println(sumSlice[int](intSlice))

// 	var float32Slice = []float32{1, 2, 3}
// 	fmt.Println(sumSlice[float32](float32Slice))

// 	fmt.Printf("Is float32 slice empty? %v", isEmpty(float32Slice))
// }

// func sumSlice[T int | float32 | float64](slice []T) T {
// 	var sum T
// 	for _, v := range slice{
// 		sum += v
// 	}
// 	return sum
// }

// func isEmpty[T any](slice []T) bool{
// 	return len(slice) == 0
// }




type ElectricEngine struct{
	kwh float32
	mpkwh float32
}

type GasEngine struct{
	gallons float32
	mpg float32
}

type Car [T GasEngine | ElectricEngine] struct{
	carMake string
	carModel string
	engine T
}


func main() {
	var gasCar = Car[GasEngine]{
		carMake: "Hyundai",
		carModel: "Elantra",
		engine: GasEngine{
			gallons: 13.2,
			mpg: 38,
		},
	}
	
	var electricCar = Car[ElectricEngine]{
		carMake: "Tesla",
		carModel: "Model 4.5",
		engine: ElectricEngine{
			kwh: 47.8,
			mpkwh: 5.78,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}