// For each independent client
package main

import "github.com/gorilla/websocket"

type Client struct{
	connection *websocket.C
}