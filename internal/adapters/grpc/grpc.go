package grpc

import (
	"github.com/antibomberman/mego-auth/internal/services"
	pb "github.com/antibomberman/mego-protos/gen/go/auth"
	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedAuthServiceServer

	service services.AuthService
}

func Register(gRPC *grpc.Server, service services.AuthService) {
	pb.RegisterAuthServiceServer(gRPC, &serverAPI{
		service: service,
	})

}
