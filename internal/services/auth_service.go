package services

import (
	"context"
	"errors"
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
	Logout(token string) error
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
		log.Println("Create User")
		userDetail, err = s.userClient.Create(context.Background(), &user.CreateUserRequest{Email: email})
		if err != nil {
			return "", fmt.Errorf("error creating user: %v", err)
		}
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
		return err
	}
	code := generateRandomCode(9999, 1000)
	//TODO send email with code
	//TODO send code 1234
	return s.redis.Set(context.Background(), email, code, time.Minute*5).Err()
}
func (s *authService) Check(token string) (bool, error) {
	blacklisted, err := s.checkIfTokenBlacklisted(token)
	if err != nil || blacklisted {
		return false, errors.New("token blacklisted")
	}
	return s.secure.Check(token)
}
func (s *authService) Parse(token string) (string, error) {
	blacklisted, err := s.checkIfTokenBlacklisted(token)
	if err != nil || blacklisted {
		return "", errors.New("token blacklisted")
	}
	claims, err := s.secure.Parse(token)
	if err != nil {
		return "", err
	}
	return claims.UserId, nil
}
func (s *authService) Logout(token string) error {
	return s.redis.Set(context.Background(), "blacklist:"+token, true, time.Hour*24).Err()
}
func (s *authService) checkIfTokenBlacklisted(token string) (bool, error) {
	exists, err := s.redis.Exists(context.Background(), "blacklist:"+token).Result()
	if err != nil {
		return false, errors.New("error checking blacklist")
	}
	return exists == 1, err
}
func generateRandomCode(max, min int) string {
	//return strconv.Itoa(rand.IntN(max-min) + min)
	return "1234"
}
