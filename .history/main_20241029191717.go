package main

import (
	"fmt"
	"net/http"
)

func setUpAPI(){
	fmt.Pr
	http.Handle("/", http.FileServer("./fronted"))
}

func main() {
	setUpAPI()
}
