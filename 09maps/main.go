package main

import "fmt"

func main() {
	fmt.Println("Welcome to map")
   //var colors map[string]string
	colors := make(map[string]string)

	colors["R"] = "Red"
	colors["Y"] = "Yellow"
	colors["O"] = "Orange"
	colors["G"] = "Green"

	fmt.Println(" List of all colors: ",colors)
	fmt.Println(" R shorts for: ",colors["R"])
	delete(colors,"O")
	
	fmt.Println(" List of all colors: ",colors)

	//looping through map
	for color,value := range colors{
		fmt.Printf("Color %v shorts for %v\n",color,value)
	}
 }