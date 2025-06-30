package main

import (
	"log"
	"net"

	authpb "github.com/DarkXPixel/Vibe/proto/auth"
	"github.com/DarkXPixel/Vibe/services/auth-service/internal"
	"google.golang.org/grpc"
)

func main() {
	cfg := internal.LoadConfig()
	err := internal.RunMigrate(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Migration is NOT OK: %v", err)
	}
	db, err := internal.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authServer := internal.NewAuthServer(db, cfg)

	authpb.RegisterAuthServiceServer(grpcServer, authServer)

	log.Println("Auth service listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
