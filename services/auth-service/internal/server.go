package internal

import (
	"context"
	"log"

	authpb "github.com/DarkXPixel/Vibe/proto/auth"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.AuthResponse, error) {
	log.Printf("Register request: username=%s, email=%s", req.GetUsername(), req.GetEmail())

	return &authpb.AuthResponse{
		UserId: "user-id-placeholder",
		Token:  "token-placeholder",
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.AuthResponse, error) {
	log.Printf("Loging request: username=%s", req.GetUsername())

	return &authpb.AuthResponse{
		UserId: "user-id-placeholder",
		Token:  "token-placeholder",
	}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *authpb.TokenRequest) (*authpb.UserSession, error) {
	log.Printf("ValidateToken request: token=%s", req.GetToken())

	return &authpb.UserSession{
		UserId:   "user-id-placeholder",
		Username: "demo-user",
		IsValid:  true,
	}, nil
}
