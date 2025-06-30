package internal

import (
	"context"
	"errors"
	"log"
	"time"

	authpb "github.com/DarkXPixel/Vibe/proto/auth"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var validate = validator.New()

type RegisterInput struct {
	Username string `validate:"required,min=3,max=20,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func fromProto(req *authpb.RegisterRequest) RegisterInput {
	return RegisterInput{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	db  DBPool
	cfg *Config
}

func NewAuthServer(db DBPool, cfg *Config) *AuthServer {
	return &AuthServer{db: db, cfg: cfg}
}

func (s *AuthServer) generateJWT(userID, username string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(s.cfg.JWTExpirationHourse)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.AuthResponse, error) {
	log.Printf("Register request: username=%s, email=%s", req.GetUsername(), req.GetEmail())

	if req.GetUsername() == "" || req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "username, email or password is null")
	}

	input := fromProto(req)
	if err := validate.Struct(input); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	var userID string
	err = s.db.QueryRow(ctx, `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`, req.Username, req.Email, string(hashedPassword)).Scan(&userID)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to insert user: %v", err)
	}

	token, err := s.generateJWT(userID, req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &authpb.AuthResponse{
		UserId: userID,
		Token:  token,
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.AuthResponse, error) {
	log.Printf("Loging request: username=%s", req.GetUsername())

	if req.GetUsername() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "username and password are required")
	}

	var (
		userID       string
		username     string
		passwordHash string
	)

	err := s.db.QueryRow(ctx, `
		SELECT id, username, password_hash
		FROM users
		WHERE username=$1
	`, req.GetUsername()).Scan(&userID, &username, &passwordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Printf("DB error: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.GetPassword()))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := s.generateJWT(userID, req.GetUsername())
	if err != nil {
		log.Printf("JWT generation error: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &authpb.AuthResponse{
		UserId: userID,
		Token:  token,
	}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *authpb.TokenRequest) (*authpb.UserSession, error) {
	log.Printf("ValidateToken request: token=%s", req.GetToken())

	tokenStr := req.GetToken()
	if tokenStr == "" {
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Error(codes.Unauthenticated, "unexcept signing method")
		}
		return s.cfg.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(codes.Internal, "invalid claims structure")
	}

	userID, _ := claims["sub"].(string)
	username, _ := claims["username"].(string)

	return &authpb.UserSession{
		UserId:   userID,
		Username: username,
		IsValid:  true,
	}, nil
}
