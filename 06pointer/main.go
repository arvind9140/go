package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int // declare a pointer

	// fmt.Println("value of ptr  is", ptr)                     // nil means it
	 myNumber := 23
	 var ptr = &myNumber

	 fmt.Println("Value of actual ptr  is", ptr)
	 fmt.Println("Value of actual ptr  is", *ptr)

	 *ptr = *ptr  * 2
	 fmt.Println("new value of myNumber  is", myNumber)

}