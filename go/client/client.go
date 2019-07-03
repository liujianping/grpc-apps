package main

import (
	pb "github.com/liujianping/grpc-apps/go/hello"
)

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var err error

	// connect to server
	conn, err := grpc.Dial("localhost:8833", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer conn.Close()

	// create client
	client := pb.NewHelloServiceClient(conn)

	// create request
	req := &pb.HelloRequest{Greeting: "JayL"}

	// call method
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// handle response
	fmt.Printf("Received: \"%s\"\n", res.Reply)
}
