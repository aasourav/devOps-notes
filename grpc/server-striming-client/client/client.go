package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

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

	stream, err := client.ServerReply(context.TODO(), &proto.HelloRequest{SomeString: "say something"})

	if err != nil {
		fmt.Println("something error")
		return
	}

	count := 0
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("user message:- ", message)
		time.Sleep(1 * time.Second)
		count++
	}

	c.JSON(http.StatusOK, gin.H{
		"message_count": count,
	})
}
