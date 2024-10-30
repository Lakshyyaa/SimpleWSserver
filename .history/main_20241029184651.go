package main

import (
	"fmt"
)

type Server struct {
	connections map[*websocket.Conn]bool
}

func newServer() *Server {
	return &Server{
		connections: make(map[*websocket.Conn]bool),
	}
}

func main() {
	fmt.Println("lf go my n")
}
