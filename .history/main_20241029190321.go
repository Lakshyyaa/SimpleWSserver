package main

import (
	"fmt"
	"net/http"
)

func setUpAPI(){
	http.Handle("/", http.FileServer("./fronted"))
}

func main() {
	setUpAPI()
}
