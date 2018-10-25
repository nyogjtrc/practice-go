package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/nyogjtrc/practice-go/version/proto"
	"github.com/nyogjtrc/practice-go/version/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

var (
	reg         = prometheus.NewRegistry()
	grpcMetrics = grpc_prometheus.NewServerMetrics()
)

func init() {
	reg.MustRegister(grpcMetrics)
}

func startServer() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor))
	opts = append(opts, grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor))
	s := grpc.NewServer(opts...)
	proto.RegisterVersionServiceServer(s, &VersionServer{})

	grpc_prometheus.Register(s)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(":9092", nil)
	}()
	s.Serve(lis)
}
