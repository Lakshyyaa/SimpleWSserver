package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	client ClientList
	
}

// This way of defining a function is to tell that this is a part of the Manager struct but defined outside
// and this initial pointer called the receiver says this function is accessed by this only
func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection!")
	conn, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	log.Println("closed connection!")
}

func NewManager() *Manager {
	return &Manager{}
}
