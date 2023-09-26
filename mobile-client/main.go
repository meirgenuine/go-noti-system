package main

import (
	"context"
	"log"
	"net/url"
	"time"

	pb "github.com/meirgenuine/go-noti-system/grpc-server/grpc-server/server"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

func main() {
	// gRPC setup
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewNotiServiceClient(conn)

	// WebSocket setup
	u := url.URL{Scheme: "ws", Host: "localhost:8568", Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("WebSocket dial error: %v", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}
		log.Printf("Received notification: %s", message)

		// Send the message to gRPC server
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err = grpcClient.GetNoti(ctx, &pb.NotiRequest{Notification: string(message)})
		if err != nil {
			log.Fatalf("Could not send notification to gRPC server: %v", err)
		}
	}
}
