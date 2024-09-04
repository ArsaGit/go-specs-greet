package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpcserver "github.com/ArsaGit/go-specs-greet/adapters/grpcserver"
	"github.com/ArsaGit/go-specs-greet/domain/interactions"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreetServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type GreetServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (g GreetServer) Curse(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	/* TODO Extract domain logic
	Use specification as a unit test against your domain logic
	*/
	return &grpcserver.GreetReply{Message: fmt.Sprintf("Go to hell, %s!", request.Name)}, nil
}

func (g GreetServer) Greet(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{Message: interactions.Greet(request.Name)}, nil
}
