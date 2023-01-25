package main

import (
	"context"
	v1 "github.com/songtomtom/go-grpc/proto/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	v1.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
