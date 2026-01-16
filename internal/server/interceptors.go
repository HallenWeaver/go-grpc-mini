package server

import (
	"context"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var unauthenticatedMethods = map[string]bool{
	"/user.v1.UserService/CreateUser": true,
}

func LoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	resp, err := handler(ctx, req)

	log.Printf("method=%s, duration=%s, err=%v", info.FullMethod, time.Since(start), err)

	return resp, err
}

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if unauthenticatedMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeaders := first(md["authorization"])
	if authHeaders == "" {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}

	userID, err := validateBearerToken(authHeaders)
	if err != nil {
		return nil, err
	}

	ctx = WithUserID(ctx, userID)

	return handler(ctx, req)
}

func validateBearerToken(authHeader string) (string, error) {
	const prefix = "Bearer "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		return "", status.Error(codes.Unauthenticated, "authorization must use Bearer scheme")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	if token == "" {
		return "", status.Error(codes.Unauthenticated, "empty bearer token")
	}

	// Demo Policy: token == userID
	return token, nil
}

func first(values []string) string {
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
