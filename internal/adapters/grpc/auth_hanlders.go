package grpc

import (
	"context"
	"fmt"
	pb "github.com/antibomberman/mego-protos/gen/go/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s serverAPI) LoginByEmail(ctx context.Context, req *pb.LoginByEmailRequest) (*pb.LoginByEmailResponse, error) {
	fmt.Println("LoginByEmail", req.Email, req.Code)
	token, err := s.service.LoginByEmail(req.Email, req.Code)

	if err != nil {
		return nil, err
	}

	return &pb.LoginByEmailResponse{
		Token: token,
	}, err

}
func (s serverAPI) LoginByEmailSendCode(ctx context.Context, req *pb.LoginByEmailSendCodeRequest) (*pb.LoginByEmailSendCodeResponse, error) {
	log.Println("LoginByEmailSendCode", req.Email)
	err := s.service.LoginByEmailSendCode(req.Email)
	if err != nil {
		log.Println("Error sending code", err)
		return nil, err
	}
	log.Println("Code sent successfully", req.Email)
	return &pb.LoginByEmailSendCodeResponse{
		Success: true,
	}, nil
}
func (s serverAPI) Parse(ctx context.Context, req *pb.ParseRequest) (*pb.ParseResponse, error) {

	userId, err := s.service.Parse(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.ParseResponse{
		UserId: userId,
	}, nil
}
func (s serverAPI) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	fmt.Println("Check", req.Token)
	check, err := s.service.Check(req.Token)
	if err != nil {
		log.Println("Error checking token", err)
		return nil, status.Errorf(codes.NotFound, "method Check not implemented")
	}
	return &pb.CheckResponse{
		Success: check,
	}, nil
}
func (s serverAPI) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := s.service.Logout(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{
		Success: true,
	}, nil
}
