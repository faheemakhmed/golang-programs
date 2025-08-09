//what is the difference between Printf and Println?
//Printf prints formatted output using format specifiers (no automatic newline), while
//Println prints arguments separated by spaces and automatically adds a newline.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app: ")
	fmt.Println("please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)

	//strconv is a packeage used for conversion

	//the line below gives an error because it has trailing characters \r\n
	//numRating, err := strconv.ParseFloat(input, 64)
	//hence we use " strings "( it is another package)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your numRating: ", numRating+1)
	}
}
