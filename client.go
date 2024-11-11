// For each independent client
package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	// to avoid concurrent write on the websocket
	egress chan Event
	// earlier it was []byte but later changed to Event as now eahc message sent
	// will be in an event wrapper
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
	}
}

// Reads message it receives from the client frontend
func (c *Client) ReadMessages() {
	defer func() {
		// cleanup function to remove the client
		c.manager.removeClient(c)
	}()
	for {
		_, payLoad, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Fatal(err)
			}
			break
		}
		// this unmarshalls the payload received which is an object of event class defined in the
		// fronted and is put in a request struct and routeEvent checks event type does
		// what it shoudl using the type and its respective eventHandler
		var request Event
		er := json.Unmarshal(payLoad, &request)
		if er != nil {
			log.Println("error marshalling event: ", er)
			break
		}
		errr := c.manager.routeEvent(request, c)
		if errr != nil {
			log.Println("error in routing event")
		}

	}
}

// Writes message it receives to the client frontend
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
			data, er := json.Marshal(messages)
			if er != nil {	
				log.Println(er)
			}
			err := c.connection.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("failed to send message ", err)
			}
			log.Println("message is sent")
		}
	}
}
