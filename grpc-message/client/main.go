package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/nyogjtrc/practice-go/grpc-message/message"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	for i := 0; i < 30; i++ {
		sent(conn)
	}
}

func sent(conn *grpc.ClientConn) {
	c := pb.NewMessagerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, _ := c.Echo(ctx, &pb.Request{Message: "hello"})
	fmt.Println(r.Message)
}
