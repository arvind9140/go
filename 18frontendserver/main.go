package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	performGetRequest()
	
}

func performGetRequest(){
	const myurl = "http://localhost:8000/v1/api/users/getdata?userId=65c0b29f1564e92b81703b4d"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
		fmt.Println("Status Code : ", response.Status)
		fmt.Println("Content Length : ", response.ContentLength)

		content, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Content : ", string(content))
		

	
	
	
}