package main

import (
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocketServer(t *testing.T) {
	go main()

	time.Sleep(time.Second * 2)

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8568/ws", nil)
	if err != nil {
		t.Fatalf("Failed to connect to the WebSocket server: %v", err)
	}
	defer c.Close()

	_, message, err := c.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read from the WebSocket server: %v", err)
	}

	expectedPrefix := "Notification!"
	if string(message)[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("Expected message to start with %s, but got %s", expectedPrefix, message)
	}
}
