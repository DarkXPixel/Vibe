package main

import (
	"log"
	"net"

	authpb "github.com/DarkXPixel/Vibe/proto/auth"
	"github.com/DarkXPixel/Vibe/services/auth-service/internal"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authServer := internal.NewAuthServer()

	authpb.RegisterAuthServiceServer(grpcServer, authServer)

	log.Println("Auth service listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
