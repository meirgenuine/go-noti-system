package main

import (
	"context"
	"log"
	"net"

	pb "github.com/meirgenuine/go-noti-system/grpc-server/server"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedNotiServiceServer
}

func (s *server) GetNoti(ctx context.Context, in *pb.Noti) (*emptypb.Empty, error) {
	// Context cancellation
	select {
	case <-ctx.Done():
		log.Printf("Context canceled: %v", ctx.Err())
		return nil, ctx.Err()
	default:
		log.Printf("Received Notification: %v", in.GetNotification())
		return &emptypb.Empty{}, nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterNotiServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
