package main

import (
	"errors"
	"fmt"
	"io"
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

func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.HelloResponse{Reply: fmt.Sprintf("message: %v from server", i)})
		if err != nil {
			return errors.New("unable to send data from server")
		}
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(req.SomeString)
	}

	return nil
}
