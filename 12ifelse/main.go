package main

import "fmt"

func main() {
	fmt.Println("this is conditional statements in golang")

	LoginCount := 20
	var Result string

	if LoginCount < 10 {
		Result = "less than 10"
	} else if LoginCount > 10 {
		Result = "greater than 10"
	} else {
		Result = "exactly 10"
	}

	fmt.Println(Result)
}
