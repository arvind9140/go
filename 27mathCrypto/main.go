package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to Math")

	// var mynumberOne  int  = 2
	// var mynumberTwo float64 = 4.5
	// fmt.Println("The sum is:", mynumberOne + int(mynumberTwo))

	//random number

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+1)

	// random from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5)) // returns a random integer between 0 and  5
	fmt.Println(myRandomNumber)

}