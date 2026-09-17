package service

import (
	"quickdeal/dto"
	"quickdeal/entity"
	"quickdeal/repository"

	"github.com/google/uuid"
)

type AdService struct {
	adRepo   *repository.AdGormRepository
	userRepo *repository.UserGormRepository
}

func NewAdService(adRepo *repository.AdGormRepository, userRepo *repository.UserGormRepository) *AdService {
	return &AdService{adRepo: adRepo, userRepo: userRepo}
}

func (s *AdService) Create(userID uuid.UUID, req dto.CreateAdRequest) (*dto.AdResponse, error) {
	ad := entity.Ad{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		AuthorID:    userID,
	}
	if err := s.adRepo.Create(&ad); err != nil {
		return nil, err
	}

	author, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.AdResponse{
		ID:          ad.ID,
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		Author:      author.Username,
		CreatedAt:   ad.CreatedAt,
	}, nil
}

func (s *AdService) List() ([]dto.AdResponse, error) {
	ads, err := s.adRepo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.AdResponse, 0, len(ads))
	for _, ad := range ads {
		responses = append(responses, dto.AdResponse{
			ID:          ad.ID,
			Title:       ad.Title,
			Description: ad.Description,
			Price:       ad.Price,
			Author:      ad.Author.Username,
			CreatedAt:   ad.CreatedAt,
		})
	}

	return responses, nil
}
