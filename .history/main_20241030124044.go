package main

import (
	"fmt"
	"net/http"
)

func setUpAPI() {
	fmt.Println("calling the set up API function")
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}

func main() {
	setUpAPI()
}
