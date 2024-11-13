package main

import (
	"encoding/json"
	"time"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage = "send_message"
	// for all messages that user types and sends to server
	EventNewMessage  = "new_message"
	// for all messages that are now sent to other clients from the server
	EventChangeRoom = "change_room"
	// for all messages that ask for changing room
)

// payload format for messages to server from user(client)
type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}
// payload format for messages to other clients from the server(which it receives from one user)
type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}

type ChangeRoomEvent struct{
	Name string `json:"name"`
}