package main

import (
	"context"

	proto "grpc.dev/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewExampleClient(conn)
	req := &proto.HelloRequest{SomeString: "hello from client"}

	client.ServerReply(context.TODO(), req)
}
