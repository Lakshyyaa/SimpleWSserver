// For each independent client
package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}

func (c *Client) ReadMessages() {
	defer func(){
		// cleanup function to remove the client
		c.manager.removeClient(c)
	}()
	for {
		msgType, payLoad, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatal(err)
			}
			break
		}
		log.Println(msgType)
		log.Println(string(payLoad))
	}
}
