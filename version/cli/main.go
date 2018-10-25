package main

import (
	"context"
	"fmt"

	"github.com/nyogjtrc/practice-go/version/proto"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(":8888", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	getServerVersion(conn)
}

func getServerVersion(conn *grpc.ClientConn) {
	client := proto.NewVersionServiceClient(conn)
	reply, err := client.Version(context.Background(), &proto.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server:\n Version: %s\n Git Commit: %s\n Build Time: %s\n",
		reply.GetVersion(),
		reply.GetGitCommit(),
		reply.GetBuildTime(),
	)
}
