package main

import "fmt"

func main() {
	fmt.Println("array in golang")

	//it necessary to provide the no. elements inside an array , during declaration
	var FruitList [4]string

	FruitList[0] = "apple"
	FruitList[1] = "watermelon"
	FruitList[2] = "banana"
	FruitList[3] = "mango"

	fmt.Println("fruit list is : ", FruitList)
	fmt.Println("length of fruit list is : ", len(FruitList))

	var VegList = [3]string{"potato", "mushroom", "radish"}
	fmt.Println("veggie lst is : ", len(VegList))

	//theres nothing much in array in golang
	//no searching etc.
}
