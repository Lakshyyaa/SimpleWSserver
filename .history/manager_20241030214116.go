package main

import "github.com/gorilla/websocket"

var (
	webSocketUpgrader = websocket.Upgrader{
		 ReadBufferSize:=1024 
		 WriteBufferSize:=1024
	}
)