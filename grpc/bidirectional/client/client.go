package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

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

	stream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("Something error")
		return
	}

	send, recieve := 0, 0
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.HelloRequest{SomeString: "message " + strconv.Itoa(i) + " from client"})
		if err != nil {
			fmt.Println("unable to send data")
			return
		}
		send++

	}

	if err := stream.CloseSend(); err != nil {
		log.Println(err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("Server message:- ", message)
		recieve++
	}

	c.JSON(http.StatusOK, gin.H{
		"message_sent":    send,
		"message_recieve": recieve,
	})
}
