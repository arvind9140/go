package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("If-else in Golang")
// 	loginCount := 10
	
// 	var result string

// 	if loginCount < 10 {
// 		result = "Regular User"
// 	}else if loginCount>10 {
// result = "Frequent User" 
// 	} else 
// 	{
// 		result = "Something Else"
// 	}


// fmt.Println("User is a ", result)

// switch statement

// rand.Seed(time.Now().UnixNano())
// diceNumber := rand.Intn(6) + 1
// fmt.Println("Dice Number is ", diceNumber)

// switch diceNumber{
// 	case 1:
// 		fmt.Println("You rolled dice and got 1")
// 	case 2:
// 		fmt.Println("You rolled dice and got 2")
// 	case 3:
// 		fmt.Println("You rolled dice and got 3")
// 	case 4:
// 		fmt.Println("You rolled dice and got 4")
// 	case 5:
// 		fmt.Println("You rolled dice and got 5")
// 	case 6:
// 		fmt.Println("You rolled dice and got 6")
// 	default:
// 		fmt.Println("Something went wrong")


// }

days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"} 
fmt.Println(days)
// for d:=0; d<len(days);d++{
// 	fmt.Print(days[d])

// }
// for i :=range days{
// fmt.Println(days[i])
// }

// for _, day:=range days{
// 	fmt.Println("index is  and value is %v \n", day)
// }
rougueValue  := 1
for rougueValue <= 10{
	if rougueValue ==2{
		goto lco
	}
	if rougueValue == 5{
		rougueValue++
	    continue
	}
	
	fmt.Println("Value is :" , rougueValue)
	rougueValue++
    
}
lco:
    fmt.Println("This is lco")

}

