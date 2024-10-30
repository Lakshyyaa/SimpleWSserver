package main

import "github.com/gorilla/websocket"

var (
	webSocketUpgrader = websocket.Upgrader{
		 ReadBufferSize, WriteBufferSize:=1024
	}
)