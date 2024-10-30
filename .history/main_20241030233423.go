package main

import (
	"fmt"
	"log"
	"net/http"
)

func setUpAPI() {
	fmt.Println("calling the set up API function")
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}

func main() {
	manager:=NewManager()
	setUpAPI()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
