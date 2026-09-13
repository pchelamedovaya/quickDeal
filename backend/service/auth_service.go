package service

import (
	"errors"
	"time"

	"quickdeal/dto"
	"quickdeal/entity"
	"quickdeal/repository"
	"quickdeal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrEmailTaken          = errors.New("email already taken")
	ErrUsernameTaken       = errors.New("username already taken")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type AuthService struct {
	userRepo         *repository.UserGormRepository
	refreshTokenRepo *repository.RefreshTokenGormRepository
	accessSecret     string
	accessTTL        time.Duration
	refreshTTL       time.Duration
}

func NewAuthService(
	userRepo *repository.UserGormRepository,
	refreshTokenRepo *repository.RefreshTokenGormRepository,
	accessSecret string,
	accessTTL time.Duration,
	refreshTTL time.Duration,
) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		accessSecret:     accessSecret,
		accessTTL:        accessTTL,
		refreshTTL:       refreshTTL,
	}
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

func (s *AuthService) Login(req dto.LoginRequest) (*dto.TokenResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	return s.createTokenPair(user.ID)
}

func (s *AuthService) Refresh(req dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	tokenHash := utils.HashRefreshToken(req.RefreshToken)

	stored, err := s.refreshTokenRepo.FindByTokenHash(tokenHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidRefreshToken
		}
		return nil, err
	}

	if stored.Revoked || time.Now().After(stored.ExpiresAt) {
		return nil, ErrInvalidRefreshToken
	}

	if err := s.refreshTokenRepo.Revoke(stored.ID); err != nil {
		return nil, err
	}

	return s.createTokenPair(stored.UserID)
}

func (s *AuthService) Logout(req dto.RefreshTokenRequest) error {
	tokenHash := utils.HashRefreshToken(req.RefreshToken)

	stored, err := s.refreshTokenRepo.FindByTokenHash(tokenHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrInvalidRefreshToken
		}
		return err
	}

	return s.refreshTokenRepo.Revoke(stored.ID)
}

func (s *AuthService) createTokenPair(userID uuid.UUID) (*dto.TokenResponse, error) {
	accessToken, err := utils.GenerateAccessToken(userID, s.accessSecret, s.accessTTL)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	refreshTokenEntity := entity.RefreshToken{
		UserID:    userID,
		TokenHash: utils.HashRefreshToken(refreshToken),
		ExpiresAt: time.Now().Add(s.refreshTTL),
	}
	if err := s.refreshTokenRepo.Create(&refreshTokenEntity); err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
