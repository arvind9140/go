package main

import (
	"fmt"
	// "sort"
)

func main() {

	fmt.Println("Welcome to slices")
	// var fruits = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitlist : %T\n", fruits) //
	// fruits =  append(fruits, "Mango","Banana")
	// fmt.Println(fruits)

	// fruits  = append(fruits[:3])
	// fmt.Println(fruits)
	// //Slicing a slice
	// highScores := make([]int,4)
	// highScores[0] = 234
	// highScores[1] = 934
	// highScores[2] = 734
	// highScores[3] = 534
	// highScores = append(highScores, 555, 666, 321)
	// fmt.Println(highScores)

	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove  an element from the slice using splice concept
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...);
	fmt.Println(courses)

}