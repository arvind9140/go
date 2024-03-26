package main

import "fmt"

func main() {
	
	fmt.Println("Welcome to function")

	result := add(10, 20)
	fmt.Println(result)
	greeter()
	proRes := proAdder(2,5,7,8,33,44,5564)
	fmt.Println("pro res: ", proRes)

}

func greeter()  {
	fmt.Println("Hello, I am a Go Function!")
}

func add(valOne int, valTwo int) int{
	return valOne +valTwo
	
}
func proAdder(values ...int) int{
	total := 0
	for _, val := range values{
		total += val
	}
	return  total
	
}