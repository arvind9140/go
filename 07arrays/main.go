package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")


	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[3] = "grape"

	fmt.Println("Fruits: ", fruitList)
	fmt.Println("Fruits: ", len(fruitList))

	var vegList  = [3]string{"potato", "tomato", "brocoli" }

	fmt.Println("Vegetables: ", vegList)
	fmt.Println("Vegetables: ", len(vegList))

}