package main

import (
	pb "github.com/liujianping/grpc-apps/go/helloworld"
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
func (hs *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// create response
	res := &pb.HelloReply{
		Message: fmt.Sprintf("hello %s from go", req.Name),
	}

	return res, nil
}

func main() {
	var err error

	// create socket listener
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// create server
	helloServer := &HelloServer{}

	// register server with grpc
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, helloServer)

	log.Println("server serving at: :50051")
	// run
	s.Serve(l)
}
