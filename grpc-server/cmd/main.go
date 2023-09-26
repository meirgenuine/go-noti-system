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
	log.Printf("Received Notification: %v", in.GetNotification())
	return &emptypb.Empty{}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotiServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
