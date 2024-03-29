package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rishichawda/demo_grpc_go/proto/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type server struct {
	pb.UnimplementedHelloWorldGreeterServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
