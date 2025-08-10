package main

import "fmt"

func main() {
	fmt.Println("this is tutorial on pointers in golang")
	// var ptr *int - creating a pointer that stores an integer
	// fmt.Println("value of pointer is : ", ptr)
	// //output->value of pointer is :  <nil>

	MyNum := 32
	var ptr = &MyNum // creates a pointer (ptr) which references to MyNum

	fmt.Println("address of pointer in memory is: ", ptr) //we get the actual memory address
	fmt.Println("value of pointer is: ", *ptr)            //we get the value stored in the memory address(e.g. 32)

	//*ptr is actually 32
	*ptr = *ptr + 2
	fmt.Println("new value after addition is: ", *ptr)  //returns 34 (which is 32+2)
	fmt.Println("new value after addition is: ", MyNum) //returns 34 (which is 32+2)
}
