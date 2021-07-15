## What is gRPC?

gRPC is a Remote Procedure Call. It's a transfeer protocol (Protobuf - Protocol buffer) created by google. This protocol is based in a "contract" that must be followed by all the services that will use this feature.

The gRPC compiles the "contract" and generates the functions we need to make it work.

## http1 vs gRPC (http2)

HTTP1 is the most common transfer protocol which every developer knows it. It's easy to understand and manipulate data with it, but it has much more costs than gRPC in terms of performance and memory costs.

The difference on them is the format that the message is sent, which is plain text on http and binary on http2

When the gRPC is compiled, a bin file is generated (the bin files are lighter than http1 plain texts) and because of that, this protocol can be 10 times faster than http1 to send messages and 7 times faster to send message to other services.

## When should we use gRPC

As I mentioned above, the gRPC needs a contract that must be followed by the sender and receiver, because of that, it can be hard to implement with partners if they don't know how to use and implement it.

This scenario changes when we talk about microservices, those are internal services developed by the team and in that case, the implementation of gRPC can provide a considerable improvement on the services performance.

## Command to generate proto 
``
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
``

## Testing gRPC Server

To test the gRPC feature, we must first initialize the gRPC server. You can do that by running the following command:

``go run cmd/server/server.go``


## Data Transfer

When the gRPC server is up, open a new terminal and run the client:

``go run cmd/client/client.go``

## Data Transfer as Stream

The gRPC allow us to transfer data as stream (stream transfer happens at real time). I've implemented a simple fibonacci function that generates 20 numbers of fibonacci sequence to test this functionality of gRPC.
To run that test, open a new terminal with the gRPC server up and use the following command

``go run cmd/stream/stream.go``
