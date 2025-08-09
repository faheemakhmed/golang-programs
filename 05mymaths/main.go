package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//fmt.Println("welcome to maths in golang")

	// var numberOne int = 2
	// var numberTwo float64 = 4.5
	// fmt.Println("the sum of numberOne and numberTwo is : ", numberOne + int(numberTwo))
	//in the above approach we loose precision - we loose the 0.5 when converted to int

	//how to generate random numbers in golang?
	//we've got 2 packages - 'math/rand' and 'crypto/rand'

	//rand.Seed() -> is deprecated

	//fmt.Println(rand.Intn(500))// this syntax is for math/rand

	//random from crypto
	MyRandomnNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(MyRandomnNum)

}
