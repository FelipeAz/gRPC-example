package main

import (
	"log"
	"net"

	"github.com/FelipeAz/gRPC-example/pb/pb"
	"github.com/FelipeAz/gRPC-example/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer, &servers.Math{})
	reflection.Register(grpcServer)

	// Creates gRPC Server
	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Cannot start the gRPC server: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to access server reflection: %v", err)
	}
}