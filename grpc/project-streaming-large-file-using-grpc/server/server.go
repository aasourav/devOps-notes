package main

import (
	"io"
	"net"
	"os"

	proto "file-uplad.dev/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedStreamUploadServer
}

func main() {
	lisntener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterStreamUploadServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lisntener); e != nil {
		panic(e)
	}
}

func (s server) Upload(stream proto.StreamUpload_UploadServer) error {
	var fileBytes []byte
	var fileSize int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		chunks := req.GetChunks()
		fileBytes = append(fileBytes, chunks...)
		fileSize += int64(len(chunks))
	}

	f, err := os.Create("./abc.iso")
	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.Write(fileBytes)
	if err2 != nil {
		return err2
	}

	return stream.SendAndClose(&proto.UploadResponse{FileSize: fileSize, Message: "successfully uploaded"})
}
