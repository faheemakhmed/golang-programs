package main

import "fmt"

// outside method(main) you are not allowed to use ( := ) operator
//example-
//jwt := 3000.00
//the above is not allowed outside of main(method)
//allowed syntax is
//var jwt = 3000
//var jwt int = 3000

const LoginToken string = "abcdefgh" //public because the first letter of constant is uppercase

func main() {
	var username string = "fahim"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 256.4654895641654
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T \n", anothervariable)

	//implicit type
	var website = "youtube.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)

	//no var style
	// := (walrus operator)
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	//constant
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
