//here, we learn about -> comma ok syntax ,error ok syntax
//taking user input
//bufio package , os package

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to my website"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //uppercase characters define which one is public : remember
	fmt.Println("enter the rating for our pizza: ")

	//comma ok || comma err -> input, _ {input is the expected(kinda like "try"); _ is in case of error(_ is kinda like a "catch")}
	//comma ok || comma err -> input,_ || _,err || input,err (there are better conventions to write this)

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating ", input)
	fmt.Printf("type of this rating is: %T \n", input)

}
