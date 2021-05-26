package main

import (
	"context"
	"io"
	"log"

	"github.com/FelipeAz/gRPC-example/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
		return
	}
	defer connection.Close()

	client := pb.NewMathServiceClient(connection)
	request := &pb.FibonacciRequest{
		N: 20,
	}

	responseStream, err := client.Fibonacci(context.Background(), request)
	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("Fibonacci: %v", stream.GetResult())
	}
}
