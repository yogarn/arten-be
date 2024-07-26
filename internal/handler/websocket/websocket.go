package websocket

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
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
	clients   map[uuid.UUID]map[*websocket.Conn]bool
	broadcast map[uuid.UUID]chan []byte
	mu        sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:   make(map[uuid.UUID]map[*websocket.Conn]bool),
		broadcast: make(map[uuid.UUID]chan []byte),
	}
}

func (manager *WebSocketManager) HandleConnections(chatId uuid.UUID, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	manager.mu.Lock()
	if _, ok := manager.clients[chatId]; !ok {
		manager.clients[chatId] = make(map[*websocket.Conn]bool)
		manager.broadcast[chatId] = make(chan []byte)
		go manager.HandleMessages(chatId)
	}
	manager.clients[chatId][conn] = true
	manager.mu.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			manager.mu.Lock()
			delete(manager.clients[chatId], conn)
			manager.mu.Unlock()
			break
		}
		manager.Broadcast(chatId, msg)
	}
}

func (manager *WebSocketManager) HandleMessages(chatId uuid.UUID) {
	for {
		msg := <-manager.broadcast[chatId]

		manager.mu.Lock()
		for client := range manager.clients[chatId] {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(manager.clients[chatId], client)
			}
		}
		manager.mu.Unlock()
	}
}

func (manager *WebSocketManager) Broadcast(chatId uuid.UUID, msg []byte) {
	manager.broadcast[chatId] <- msg
}
