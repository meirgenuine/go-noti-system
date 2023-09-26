package main

import (
	"context"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	pb "github.com/meirgenuine/go-noti-system/grpc-server/server"
	"google.golang.org/grpc"
)

func main() {
	// gRPC setup
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
		return
	}
	defer conn.Close()

	grpcClient := pb.NewNotiServiceClient(conn)

	// WebSocket setup
	u := url.URL{Scheme: "ws", Host: "localhost:8568", Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("WebSocket dial error: %v", err)
		return
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			return
		}

		log.Printf("Received notification: %s", message)

		go func(msg []byte) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			_, err = grpcClient.GetNoti(ctx, &pb.Noti{Notification: string(msg)})
			if err != nil {
				log.Printf("Could not send notification to gRPC server: %v", err)
				return
			}

			log.Printf("Sent notification to gRPC server")
		}(message)
	}
}
