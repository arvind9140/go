package main

import "fmt"

func main() {
	var username string = "Arvind"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)


	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.21345678888888765432
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)


	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "Arvind.com"
	fmt.Println("Website : ",website)

	// no var style

	numberOfUser:= 300000
	fmt.Println(numberOfUser)
}