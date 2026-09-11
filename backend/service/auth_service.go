package service

import (
	"errors"

	"quickdeal/dto"
	"quickdeal/entity"
	"quickdeal/repository"
	"quickdeal/utils"

	"gorm.io/gorm"
)

var (
	ErrEmailTaken    = errors.New("email already taken")
	ErrUsernameTaken = errors.New("username already taken")
)

type AuthService struct {
	userRepo *repository.UserGormRepository
}

func NewAuthService(userRepo *repository.UserGormRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
		return nil, ErrEmailTaken
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if _, err := s.userRepo.FindByUsername(req.Username); err == nil {
		return nil, ErrUsernameTaken
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: passwordHash,
	}
	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
