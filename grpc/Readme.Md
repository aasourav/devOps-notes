
## Types of gRpc:

1. Unary operation 
    - Client  ---req---> <--res--- Server
2. Server streaming
    - Server to Client data stream
3. Client striming
    - Client to Server stream
4. Bidirectional streaming
    - Client to Server , Server to Client streaming

## What else does gRPC offer?
 - Metadata (headers)
 - Streaming

```proto
syntax = "proto3";


option go_package = "./"; // where auto generate code will be generated

// Here message is a proto keyword
// HelloReqeust is a message name
// inside HelloRequest SomeString is string type
// SomeString = 1 means it's sequence number 
message HelloRequest{
    string SomeString = 1;
}

message HelloResponse{
    string Reply = 1;
}

// service is a proto keyword
service Example{
    //rpc fn
    rpc ServerReply(HelloRequest) returns (HelloResponse);
}



```

To run this proto file it requires protobuf compiler and package which language you want to convert . in my case it's go

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# export PATH="$PATH:$(go env GOPATH)/bin" #(Optional, it should automaticlly set)
sudo apt install -y protobuf-compiler
protobuf --version
protoc --go-grpc_out=> --go_out=. *.proto
```

after running this command two file will be generated. one is `exaample_grpc.pb.go` another one is `example.pb.go`
all models  ( message HellResponse ) will be put in `example.pb.go` and all the services ( service Example ) will be inside `example_grpc.pb.go`

## Unary operation in gRPC

initial forlder structre 
```md
- client
- protoc
- server
```


## Source: [gRPC Playlist](https://www.youtube.com/playlist?list=PL8fnAiiuQeFspXaocqil47hesUod0y85C)