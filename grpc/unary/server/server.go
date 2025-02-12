package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	proto "grpc.dev/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer // this should be provide here. see this keyword `UnimplementedExampleServer` on example_grpc.pb.go. this fn provide services fn
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer()                     //engine
	proto.RegisterExampleServer(srv, &server{}) // here we register the svc `Example` service
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("recieve request from client", req.SomeString)
	fmt.Println("hello from server")
	return &proto.HelloResponse{}, errors.New("")
}
