package main

import (
	"fmt"
	"time"
)

func main() {

	PresentTime := time.Now()
	fmt.Println(PresentTime)
	//remember this exactly -> 01-02-2006
	// also to print day, always send "Monday" {not monday, or any other day}
	// to get time -> "15:04:05" - exactly this
	//go to documentation in examples
	fmt.Println(PresentTime.Format("01-02-2006"))
	fmt.Println(PresentTime.Format("01-02-2006 Monday"))

	//THERE IS A LOT IN TIME(in golang)
	//EXPLORE DOCUMENTATION TO LEARN MORE
}
