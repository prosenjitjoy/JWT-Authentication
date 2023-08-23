package user

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=5,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=24"`
}

type SignUpResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=24"`
}

type SignInResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token,omitempty"`
}

type Adapter interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	Register(c context.Context, req *SignUpRequest) (*SignUpResponse, error)
	Login(c context.Context, req *SignInRequest) (*SignInResponse, error)
}
