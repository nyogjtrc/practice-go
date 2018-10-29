package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/nyogjtrc/practice-go/grpc-message/message"
	"google.golang.org/grpc"
)

type server struct {
	count int
}

func (s *server) Echo(c context.Context, req *pb.Request) (*pb.Reply, error) {
	s.count++
	fmt.Println(s.count, req.String(), time.Now())

	return &pb.Reply{
		Message: "echo: " + req.Message,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterMessagerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
