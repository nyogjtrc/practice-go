package main

import (
	"context"
	"net"

	"github.com/nyogjtrc/practice-go/version/proto"
	"github.com/nyogjtrc/practice-go/version/version"
	"google.golang.org/grpc"
)

type VersionServer struct {
}

func (s *VersionServer) Version(context.Context, *proto.Empty) (*proto.VersionMessage, error) {
	reply := new(proto.VersionMessage)
	reply.Version = version.Version
	reply.GitCommit = version.GitCommit
	reply.BuildTime = version.BuildTime
	return reply, nil
}

func startServer() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	proto.RegisterVersionServiceServer(s, &VersionServer{})
	s.Serve(lis)
}
