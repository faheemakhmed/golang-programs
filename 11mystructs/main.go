package main

import "fmt"

func main() {
	fmt.Println("structs in golang")

	Fahim := User{"Fahim", "fahim@xyz.com", true, 20}
	fmt.Println(Fahim)
	fmt.Printf("fahims details are: %+v \n", Fahim) //gives more details
	fmt.Printf("name is %v , email is %v", Fahim.Name, Fahim.Email)

}

//theres no concept of inheritance in golang or super or parent
//defining a struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
