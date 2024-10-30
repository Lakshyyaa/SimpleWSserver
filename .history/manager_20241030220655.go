package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection!")
	conn, err:=webSocketUpgrader.(w,r,nil)
}
