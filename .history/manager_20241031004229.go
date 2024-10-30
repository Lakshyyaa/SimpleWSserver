package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients ClientList
	// adding a mutex as we might have many clients connecting to the API concurrently
	sync.RWMutex
}

// This way of defining a function is to tell that this is a part of the Manager struct but defined outside
// and this initial pointer called the receiver says this function is accessed by this only
func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection!")
	conn, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	client:=NewClient(conn, m)
	m.addClient()
}

func (m *Manager) addClient(client *Client){
	m.Lock()
	if ),err:=m.cli
	defer m.Unlock()
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}
