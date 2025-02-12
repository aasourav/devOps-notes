package main

import (
	"fmt"
	"net"
	"time"

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

func (s *server) ServerReply(req *proto.HelloRequest, stream proto.Example_ServerReplyServer) error {
	fmt.Println(req.SomeString)
	time.Sleep(5 * time.Second)

	reply := []*proto.HelloResponse{
		{Reply: "Hi 1"},
		{Reply: "Hi 2"},
		{Reply: "Hi 3"},
		{Reply: "Hi 4"},
		{Reply: "Hi 5"},
		{Reply: "Hi 6"},
	}

	for _, msg := range reply {
		err := stream.SendMsg(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
