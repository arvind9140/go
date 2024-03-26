package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arvind910/momgo/routes"
)

func main() {
	fmt.Println("MongoDB  API")
	r := routes.Router()
	fmt.Println("Server is getting started.....")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server started on port 8080...")

}