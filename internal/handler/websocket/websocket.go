package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketManager struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (manager *WebSocketManager) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	manager.clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(manager.clients, conn)
			break
		}
	}
}

func (manager *WebSocketManager) HandleMessages() {
	for {
		msg := <-manager.broadcast

		for client := range manager.clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(manager.clients, client)
			}
		}
	}
}

func (manager *WebSocketManager) Broadcast(msg []byte) {
	manager.broadcast <- msg
}
