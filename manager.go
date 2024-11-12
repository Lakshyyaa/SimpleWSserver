package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

type Manager struct {
	clients ClientList
	// adding a mutex as we might have many clients connecting to the API concurrently
	sync.RWMutex
	handlers map[string]EventHandler
	otps     RetentionMap
}

// This way of defining a function is to tell that this is a part of the Manager struct but defined outside
// and this initial pointer called the receiver says this function is accessed by this only
func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection!")
	otp := r.URL.Query().Get("otp")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !m.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	conn, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := NewClient(conn, m)
	m.addClient(client)
	go client.ReadMessages()
	go client.WriteMessages()
}

func (m *Manager) loginHandler(w http.ResponseWriter, r *http.Request) {
	type userLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req userLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (req.Username == "lakshya" && req.Password == "123") {
		type respone struct {
			OTP string `json:"otp"`
		}
		otp := m.otps.NewOTP()
		resp := respone{
			OTP: otp.Key,
		}
		// creating an otp and sending to frontend
		data, er := json.Marshal(resp)
		if er != nil {
			log.Println("login handler er: ", er)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	m.clients[client] = true
	defer m.Unlock()
}

func (m *Manager) setUpEventHandlers() {
	m.handlers[EventSendMessage] = SendMessage
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	handler, ok := m.handlers[event.Type]
	if ok {
		err := handler(event, c)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
	}
}

func SendMessage(event Event, c *Client) error {
	log.Println(event)
	return nil
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	_, ok := m.clients[client]
	if ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

// each manager has its clientlist
func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
		otps:     NewRetentionMap(ctx, 5*time.Second),
	}
	m.setUpEventHandlers()
	return m
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	switch origin {
	case "http://localhost:3000":
		return true
	default:
		return false
	}
}
