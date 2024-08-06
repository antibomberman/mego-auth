package main

import (
	"context"
	"fmt"
	adapter "github.com/antibomberman/mego-auth/internal/adapters/grpc"
	"github.com/antibomberman/mego-auth/internal/clients"
	"github.com/antibomberman/mego-auth/internal/config"
	"github.com/antibomberman/mego-auth/internal/secure"
	"github.com/antibomberman/mego-auth/internal/services"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.Load()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})
	err := rdb.Ping(context.Background()).Err()

	if err != nil {
		log.Fatal("redis connection error")
	}
	log.Println("Connected to Redis")
	userClient, err := clients.NewUserClient(cfg.UserServiceAddress)
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	scr := secure.NewSecure(cfg)

	authService := services.NewAuthService(rdb, userClient, scr)

	l, err := net.Listen("tcp", ":"+cfg.AuthServiceServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPC := grpc.NewServer()
	adapter.Register(gRPC, authService)
	log.Println("Server started on port", cfg.AuthServiceServerPort)
	if err := gRPC.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Server stopped")
}
