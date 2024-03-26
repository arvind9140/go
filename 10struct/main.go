package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	kush := User{"Kushagra", "kush@gmail.com", true, 24}
	fmt.Println(kush)
	fmt.Printf("kushagra details are: %v\n", kush)
	fmt.Printf("Name is %v  and Email is %v ", kush.Name, kush.Email)

}

type User struct{
	Name    string
	Email   string
	Status  bool
	Age     int
}