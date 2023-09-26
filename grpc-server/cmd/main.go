package main

import (
	"log"
	"net"

	pb "github.com/meirgenuine/go-noti-system/grpc-server/server/server"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetNoti(in *pb.NotiRequest, stream pb.NotiService_GetNotiServer) error {
	log.Printf("Received Notification: %v", in.GetNotification())
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotiServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
