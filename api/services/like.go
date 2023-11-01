package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type LikeService struct {
	repository repository.LikeRepository
}

func NewLikeService(repository repository.LikeRepository) LikeService {
	return LikeService{
		repository: repository,
	}
}

// CreateFeed to create the Feed
func (c LikeService) Create(Like models.Like) error {
	err := c.repository.Create(Like)
	return err
}
