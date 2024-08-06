package services

import (
	"context"
	"fmt"
	"github.com/antibomberman/mego-auth/internal/clients"
	"github.com/antibomberman/mego-auth/internal/secure"
	"github.com/antibomberman/mego-protos/gen/go/user"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type AuthService interface {
	LoginByEmail(email, code string) (string, error)
	LoginByEmailSendCode(email string) error
	Check(token string) (bool, error)
	Parse(token string) (string, error)
}

type authService struct {
	secure     secure.Secure
	redis      *redis.Client
	userClient *clients.UserClient
}

func NewAuthService(rdb *redis.Client, client *clients.UserClient, secure secure.Secure) AuthService {
	return &authService{secure: secure, redis: rdb, userClient: client}
}

func (s *authService) LoginByEmail(email, code string) (string, error) {
	userDetail, err := s.userClient.GetByEmail(context.Background(), &user.Email{Email: email})
	if err != nil {
		log.Printf("Error getting user by email: %v", err)
		return "", fmt.Errorf("invalid email")
	}

	savedCode, err := s.redis.Get(context.Background(), email).Result()
	if err != nil {
		log.Printf("Error getting email code: %v", err)
		return "", fmt.Errorf("invalid code")
	}
	if savedCode != code {
		log.Printf("Invalid code != saved code")
		return "", fmt.Errorf("invalid code")
	}

	token, err := s.secure.Generate(userDetail.Id)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", fmt.Errorf("error generating token: %v", err)
	}
	return token, nil
}

func (s *authService) LoginByEmailSendCode(email string) error {
	_, err := s.userClient.GetByEmail(context.Background(), &user.Email{Email: email})
	if err != nil {
		return fmt.Errorf("invalid email")
	}
	code := generateRandomCode(9999, 1000)
	return s.redis.Set(context.Background(), email, code, time.Minute*5).Err()
}

func (s *authService) Check(token string) (bool, error) {
	return s.secure.Check(token)
}
func (s *authService) Parse(token string) (string, error) {
	claims, err := s.secure.Parse(token)
	if err != nil {
		return "", err
	}
	return claims.UserId, nil
}

func generateRandomCode(max, min int) string {
	//return strconv.Itoa(rand.IntN(max-min) + min)
	return "1234"
}
