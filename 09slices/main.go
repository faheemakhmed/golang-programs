//slices are important in golang

package main

import (
	"fmt"
)

func main() {
	// fmt.Println("slices tutorial")

	// var FruitList = []string{"apple", "tomato", "peach"}
	// fmt.Printf("type of fruitlist is %T \n", FruitList)

	// FruitList = append(FruitList, "mango", "banana")
	// fmt.Println(FruitList)

	// FruitList = append(FruitList[1:]) //deletes stuff from the slice
	// fmt.Println(FruitList)

	// FruitList = append(FruitList[1:3]) //deletes stuff from the slice
	// fmt.Println(FruitList)

	// FruitList = append(FruitList[:3]) //deletes stuff from the slice
	// fmt.Println(FruitList)

	// HighScores := make([]int, 4)

	// HighScores[0] = 234
	// HighScores[1] = 945
	// HighScores[2] = 465
	// HighScores[3] = 867
	// //HighScores[4] = 777 //this will be error because the range is 4 as declared above

	// fmt.Println(HighScores)

	// HighScores = append(HighScores, 555, 666) //these new values will be accomodated in memory
	// fmt.Println(HighScores)

	// //there's a package called "sort"
	// sort.Ints(HighScores) //sorts all in increasing order
	// fmt.Println(HighScores)

	// //sort.IntsAreSorted(HighScores) -> tells whether ints are sorted or not (returns boolean value)

	//how to remove from slices based on index
	var Courses = []string{"javascript", "pyhton", "golang", "react", "ruby"}
	fmt.Println(Courses)
	var index int = 2
	// Courses = append(Courses[:index]) //prints from index[0] to index[2] excluding last index. i.e. excluding [2]
	// fmt.Println(Courses)

	Courses = append(Courses[:index], Courses[index+1:]...)
	fmt.Println(Courses)

}
