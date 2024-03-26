package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// performPostRequest()
	performPostFormRequest()
	
}
func performPostRequest()  {
	const myurl = "http://localhost:8000/v1/api/users/login"

	requestBody := strings.NewReader(`
	{
		"user_name": "deva",
        "password": "@meranaam@20"
	}`)
	response,err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
	    	panic(err)
	}
	defer response.Body.Close()
	content, _:= ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	
}

func performPostFormRequest()  {
	const myurl = "http://localhost:8000/v1/api/users/send/otp/"

	data :=url.Values{}
	data.Add("email","the.arvind.kumar.maurya2001@gmail.com")
	response,err := http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content, _:= ioutil.ReadAll(response.Body)
	//printing the server's response
	fmt.Println(string(content))

	
}