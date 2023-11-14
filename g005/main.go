package main

import (
	"fmt"
)

// Struct
type GasEngine struct{
	mpg uint8
	gallons uint8
	Owner
	// int
}

type ElectricEngine struct{
	mpkwh uint8
	kwh uint8
	Owner
	// int
}

type Owner struct{
	name string
}

// A Method tied to the struct GasEngine
func (engine GasEngine) milesLeft() uint8 {
	return engine.gallons * engine.mpg
}

func (engine ElectricEngine) milesLeft() uint8 {
	return engine.kwh * engine.mpkwh
}

// Interface
type Engine interface{
	milesLeft() uint8
}

func canItGo(engine Engine, miles uint8) {
	if miles <= engine.milesLeft() {
		fmt.Println("(+) Yes, we can do this!")
	} else {
		fmt.Println("(x) Nah, we need more gas!")
	}
}

func main(){
	// var myEngine GasEngine = GasEngine{15, 25, Owner{"Reward"}, 5}
	var myEngine GasEngine = GasEngine{15, 25, Owner{"Reward"}}
	myEngine.gallons = 20
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.Owner.name, myEngine.int)
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	// Anonymous Struct: They are not reusable
	var secondEngine = struct{
		mpg uint8
		gallons uint8
	}{25, 15}
	fmt.Println(secondEngine)

	var teslaEngine ElectricEngine = ElectricEngine{25, 35, Owner{"Refuge"}}

	// Struct method
	fmt.Printf("[!] Total miles left in gas tank: %v\n", myEngine.milesLeft())
	canItGo(myEngine, 44)
	// canItGo(teslaEngine, 75)
	fmt.Println(teslaEngine.milesLeft())
}