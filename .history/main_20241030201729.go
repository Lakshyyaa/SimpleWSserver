package main

import (
	"fmt"
	"log"
	"net/http"
)

func setUpAPI() {
	fmt.Println("calling the set up API function")
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	log.Fatal()
}

func main() {
	setUpAPI()
}
