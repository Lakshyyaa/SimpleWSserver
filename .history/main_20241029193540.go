package main

import (
	"fmt"
	"net/http"
)

func setUpAPI(){
	fmt.Println("calling the set up API function")
	http.Handle("/", http.FileServer(ht))
}

func main() {
	setUpAPI()
}
