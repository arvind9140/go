package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	password string
	Tags     []string
}

func main() {
	fmt.Println("Hello Json")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {

	lcoCources := []course{
		{"JS", 300, "github", "abc123", []string{"web-dev", "js"}},
		{"MERN", 500, "github", "abc123", []string{"full-stack", "mern"}},
		{"PYTHON", 450, "github", "abc123", []string{"data-science", "python"}},
		{"DSA", 500, "github", "abc123", []string{"competitive", "core-cs"}},
	}
	finalJson, err := json.MarshalIndent(lcoCources, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "MERN",
                "Price": 500,
                "website": "github",
                "Tags": [
                        "full-stack",
                        "mern"
                       ]
	}
	`)
	var lcoCource course

	 checkValid := json.Valid(jsonDataFromWeb)
	   if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCource)
		fmt.Printf("%#v\n", lcoCource) //printing the struct in Go syntax
	   }else{
		fmt.Println("JSON was not valid")
	   }

	// some cases where you just want to add data to key value pairs of a map

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData{

	    fmt.Printf("Key is %v and value is %v and Type is: %T\n", key, value, value)
	    

	}
	    
	
}
