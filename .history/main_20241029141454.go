package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Server struct {
	connections map[*websocket.Conn]bool	
}



func main() {
	fmt.Println("lf go my n")
}
