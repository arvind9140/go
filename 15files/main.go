package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to  files")
	content  :=  "This needs to go in a file - arvind9140.com"
	file, err := os.Create("file_example.txt");
	// if err != nil {
	//     panic(err)
	// }else{
		checkNikErr(err)
 length, err := io.WriteString(file, content)
 if err != nil {
     panic(err)
 }else{
	 fmt.Printf("The file was created successfully, %d bytes written.\n", length)
	 defer file.Close()
	 readFile("./file_example.txt")
 }

	// }
}
 func readFile(filename string){
	databytes,err:= ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("File read successfully \n",string(databytes) )
	}
	
 }

 func checkNikErr(err error)  {
	if  err != nil {
	    panic(err)
	}
	
 }