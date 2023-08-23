package user

import (
	"context"
	"os"
	"server/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	adapter Adapter
	timeout time.Duration
}

func NewService(adapter Adapter) Service {
	return &service{
		adapter: adapter,
		timeout: 2 * time.Second,
	}
}

func (s *service) Register(c context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	res, err := s.adapter.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &SignUpResponse{
		ID:       res.ID.String(),
		Username: res.Username,
		Email:    res.Email,
	}

	return response, nil
}

type SignedDetails struct {
	ID       string
	Username string
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *SignInRequest) (*SignInResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.adapter.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, err
	}

	claims := &SignedDetails{
		ID:       user.ID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		AccessToken: token,
	}, nil
}
