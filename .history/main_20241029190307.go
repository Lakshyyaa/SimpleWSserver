package main

import (
	"fmt"
	"net/http"
)

func setUpAPI(){
	http.Handle("/", http.F)
}

func main() {
	setUpAPI()
}
