package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "grpc.dev/proto"
)

var client proto.ExampleClient

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client = proto.NewExampleClient(conn)

	r := gin.Default()
	r.GET("/sent", clientConnectionServer)

	r.Run(":8000")

	// req := &proto.HelloRequest{SomeString: "hello from client: Shi"}

	// client.ServerReply(context.TODO(), req)
}

func clientConnectionServer(c *gin.Context) {

	req := []*proto.HelloRequest{
		{SomeString: "Req 1"},
		{SomeString: "Req 2"},
		{SomeString: "Req 3"},
		{SomeString: "Req 4"},
		{SomeString: "Req 5"},
		{SomeString: "Req 6"},
		{SomeString: "Req 7"},
	}

	stream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("something error")
		return
	}

	for _, re := range req {
		err = stream.Send(re)
		if err != nil {
			fmt.Println("request not fulfil")
			return
		}
	}

	resposne, err := stream.CloseAndRecv()
	c.JSON(http.StatusOK, gin.H{
		"message_count": resposne,
	})
}
