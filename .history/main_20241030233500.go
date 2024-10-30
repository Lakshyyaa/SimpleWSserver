package main

import (
	"fmt"
	"log"
	"net/http"
)

func setUpAPI() {
	fmt.Println("calling the set up API function")
	manager:=NewManager()
	http.Handle()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}

func main() {
	setUpAPI()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
