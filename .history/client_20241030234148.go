// For each independent client
package main

import "github.com/gorilla/websocket"

type Client struct{
	connection *websocket.Conn
	manager *Manager
}
func NewClient(conn *websocket.Conn, manager *Man){

}