package usecase

import (
	"context"
	"realtime-chatapp/internal/domain"
	"realtime-chatapp/internal/helper"

	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {
	authRepository domain.AuthRepository
}

func NewAuthUsecase(authRepository domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecaseImpl{
		authRepository: authRepository,
	}
}

func (u *AuthUsecaseImpl) Regiter(ctx context.Context, request domain.RegisterRequest) error {
	pasHass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(pasHass),
	}
	err = u.authRepository.Save(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u *AuthUsecaseImpl) Login(ctx context.Context, request domain.LoginRequest) (domain.AuthResponse, error) {
	user, err := u.authRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return domain.AuthResponse{}, errors.New("email is not valid")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return domain.AuthResponse{}, errors.New("password is not valid")
	}
	token, err := helper.GenereteToken(user.ID)
	if err != nil {
		return domain.AuthResponse{}, errors.New("failed to generate token")
	}
	return domain.ToAuthResponse(user, token), nil
}
