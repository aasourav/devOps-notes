package main

import (
	"fmt"
	"io"
	"net"
	"strconv"

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
	total := 0
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: strconv.Itoa(total),
			})
		}
		if err != nil {
			return err
		}

		total++
		fmt.Println(request)
	}
}
