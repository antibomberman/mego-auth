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
		return &pb.LoginByEmailResponse{
			Message: err.Error(),
			Token:   token,
			Success: false,
		}, err
	}

	return &pb.LoginByEmailResponse{
		Message: "Successfully logged in",
		Token:   token,
		Success: true,
	}, err

}
func (s serverAPI) LoginByEmailSendCode(ctx context.Context, req *pb.LoginByEmailSendCodeRequest) (*pb.LoginByEmailSendCodeResponse, error) {
	log.Println("LoginByEmailSendCode", req.Email)
	err := s.service.LoginByEmailSendCode(req.Email)
	if err != nil {
		log.Println("Error sending code", err)
		return &pb.LoginByEmailSendCodeResponse{
			Message: err.Error(),
			Success: false,
		}, err
	}
	log.Println("Code sent successfully", req.Email)
	return &pb.LoginByEmailSendCodeResponse{
		Message: "Code sent successfully",
		Success: true,
	}, nil
}
func (s serverAPI) Parse(ctx context.Context, req *pb.ParseRequest) (*pb.ParseResponse, error) {

	userId, err := s.service.Parse(req.Token)
	if err != nil {
		return &pb.ParseResponse{
			Success: false,
			UserId:  "",
		}, err
	}
	return &pb.ParseResponse{
		Success: true,
		UserId:  userId,
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
		return &pb.LogoutResponse{
			Message: err.Error(),
			Success: false,
		}, err
	}
	return &pb.LogoutResponse{
		Message: "Successfully logged out",
		Success: true,
	}, nil
}
