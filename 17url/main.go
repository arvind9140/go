package main

import ( "fmt"
"net/url"
)

const urlString = "https://codegeex.cn"
func main() {
	fmt.Println("Hello, url!")
	result, _ := url.Parse(urlString)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	queryParams:= result.Query()
	fmt.Printf("Params type:%T\n", queryParams)
	fmt.Println(queryParams)
	for key, value := range queryParams {
		fmt.Printf("%s:%s\n", key, value)

	}
	
}