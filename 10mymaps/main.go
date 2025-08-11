package main

import "fmt"

func main() {
	fmt.Println("maps in goalng")

	Languages := make(map[string]string) //creation of a map using make()
	//[string]string this is kinda like [key:value] pairs

	Languages["js"] = "javascript"
	Languages["rb"] = "ruby"
	Languages["py"] = "python"

	fmt.Println("list of all languages: ", Languages)
	fmt.Println("js is short for: ", Languages["js"])

	//to delete - use delete()
	delete(Languages, "rb")
	fmt.Println(Languages)

	//loops in maps in golang

	for key, value := range Languages {
		fmt.Printf("for key %v, value is %v \n", key, value)
	}

}
