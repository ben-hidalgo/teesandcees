# teesandcees

Helpful
```
grpcurl -v -plaintext -d '{"name":"bob"}' localhost:50051 helloworld.Greeter/SayHello

protoc --go_out=plugins=grpc:. helloworld/helloworld.proto
```
