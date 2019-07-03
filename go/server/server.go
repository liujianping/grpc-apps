package main

import (
	pb "github.com/liujianping/grpc-apps/go/hello"
)

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServer struct{}

// SayHello says 'hi' to the user.
func (hs *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// create response
	res := &pb.HelloResponse{
		Reply: fmt.Sprintf("hello %s from go", req.Greeting),
	}

	return res, nil
}

func main() {
	var err error

	// create socket listener
	l, err := net.Listen("tcp", ":8833")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// create server
	helloServer := &HelloServer{}

	// register server with grpc
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, helloServer)

	log.Println("server serving at: :8833")
	// run
	s.Serve(l)
}
