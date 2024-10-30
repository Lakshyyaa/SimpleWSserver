package main

import "fmt"

type Server struct{
	connections make(map[*websocket.Conn]bool)
}

func main() {
	fmt.Println("lf go my n")
}