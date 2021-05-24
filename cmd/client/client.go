package main

import (
	"context"
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
	request := &pb.NewSumRequest{
		Sum: &pb.Sum{
			A: 11,
			B: 15,
		},
	}

	res, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Could not send request: %v", err)
		return
	}

	log.Println(res.Result)
}
