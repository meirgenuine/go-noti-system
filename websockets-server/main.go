package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var clients sync.Map
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer func() {
		clients.Delete(conn)
		conn.Close()
	}()
	clients.Store(conn, true)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			return
		}
	}
}

func handleNotifications() {
	for {
		time.Sleep(5 * time.Second)
		notification := "Notification! Time: " + time.Now().String()

		clients.Range(func(k, v interface{}) bool {
			client := k.(*websocket.Conn)
			err := client.WriteMessage(websocket.TextMessage, []byte(notification))
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				clients.Delete(client)
				client.Close()
			}
			return true
		})
	}
}

func main() {
	port := "8568"

	http.HandleFunc("/ws", handleConnections)

	go handleNotifications()

	log.Printf("Server started at :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
