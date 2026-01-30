package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserResponse struct {
	ID    int    `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type AuthRepository interface {
	Save(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, id int) (User, error)
}

type AuthUsecase interface {
	Regiter(ctx context.Context, request RegisterRequest) error
	Login(ctx context.Context, request LoginRequest) (AuthResponse, error)
}

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

func ToAuthResponse(user User, token string) AuthResponse {
	return AuthResponse{
		Token: token,
		User: UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
}
