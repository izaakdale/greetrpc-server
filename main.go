package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	greet "github.com/izaakdale/greetrpc-server/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	greet.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, in *greet.GreetRequest) (*greet.GreetResponse, error) {
	fmt.Println("Greet hit")
	return &greet.GreetResponse{
		Greeting: "Hello, " + in.Name,
	}, nil
}

func main() {
	gs := grpc.NewServer()
	greet.RegisterGreetServiceServer(gs, &Server{})
	grpc_health_v1.RegisterHealthServer(gs, health.NewServer())

	reflection.Register(gs)

	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDRESS"))
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Fatal(gs.Serve(lis))
}
