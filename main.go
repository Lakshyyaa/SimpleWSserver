package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func setUpAPI() {
	fmt.Println("calling the set up API function")
	ctx := context.Background()
	manager := NewManager(ctx)
	http.HandleFunc("/ws", manager.serveWS)
	http.HandleFunc("/login", manager.loginHandler)
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}

func main() {
	setUpAPI()
	log.Fatal(http.ListenAndServeTLS(":3000", "server.crt", "server.key", nil))
}

// When login is clicked, first a post req goes to /login which checks the username and password
// and if correct, returns an otp
// Now along with otp a req is made to /ws using url params route which verifies the otp,
// makes it obosolte and upgrades connection to let user work with the websocket server
