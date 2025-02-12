package main

import (
	"context"
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
	r.GET("/sent-msg/:message", clientConnectionServer)

	r.Run(":8000")

	// req := &proto.HelloRequest{SomeString: "hello from client: Shi"}

	// client.ServerReply(context.TODO(), req)
}

func clientConnectionServer(c *gin.Context) {
	message := c.Param("message")
	req := &proto.HelloRequest{SomeString: "Hello bro " + message}
	client.ServerReply(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully: " + message,
	})
}
