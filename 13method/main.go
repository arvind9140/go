package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	kush := User{"Kushagra", "kush@gmail.com", true, 24}
	fmt.Println(kush)
	fmt.Printf("kushagra details are: %v\n", kush)
	fmt.Printf("Name is %v  and Email is %v \n", kush.Name, kush.Email)
	kush.GetStatus()
	kush.NewMail()
fmt.Printf("Name is %v  and Email is %v \n", kush.Name, kush.Email)
}

type User struct{
	Name    string
	Email   string
	Status  bool
	Age     int
}


func (u User) GetStatus()  {
	fmt.Println("Is user active: ", u.Status)
	
}
func (u User) NewMail()  {
	u.Email = "test@gmail.com"
	fmt.Println("This is new mail: ", u.Email)
}