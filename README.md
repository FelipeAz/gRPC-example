## Command to generate proto 
``
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
``

## Testing gRPC Server
