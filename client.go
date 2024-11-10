// For each independent client
package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan []byte
	// to avoid concurrent write on the websocket
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte),
	}
}

func (c *Client) ReadMessages() {
	defer func() {
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
		for wsclient := range c.manager.clients {
			wsclient.egress <- payLoad
		}
		log.Println(msgType, " is the message type")
		log.Println(string(payLoad))
	}
}

func (c *Client) WriteMessages() {
	defer func() {
		// cleanup function to remove the client
		c.manager.removeClient(c)
	}()
	for {
		select {
		// reading from the egress and sending to ws server
		case messages, ok := <-c.egress:
			// ok is a bool to tell if channel open or closed
			if !ok {
				err := c.connection.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					log.Println("connection closed, ", err)
				}
				return
			}
			err := c.connection.WriteMessage(websocket.TextMessage, messages)
			if err != nil {
				log.Println("failed to send message ", err)
			}
			log.Println("message is sent")
		}
	}
}
