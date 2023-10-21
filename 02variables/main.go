package main

import "fmt"

func main() {
	var username string = "Ahmad"
	fmt.Println(username)
	fmt.Printf("Variable is of type is: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type is: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type is: %T \n", smallVal)
}
