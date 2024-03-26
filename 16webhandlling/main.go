package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url =  "https://github.com/arvind9140"
func main() {
	fmt.Println("welcome to web handlling")
	 response, err := http.Get(url)
	 if err != nil {
	     panic(err)
	 }else{
// fmt.Println("Response Type: %T\n", response)
   defer response.Body.Close()

  databytes,err :=  ioutil.ReadAll(response.Body)
  if err !=nil{
	   panic(err)
  }else{
	content:= string(databytes)
	fmt.Println(content)
  }

	 }

}