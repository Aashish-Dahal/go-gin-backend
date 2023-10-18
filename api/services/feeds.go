package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type FeedService struct {
	repository repository.FeedRepository
}

func NewFeedService(repository repository.FeedRepository) FeedService {
	return FeedService{
		repository: repository,
	}
}

// CreateFeed to create the Feed
func (c FeedService) CreateFeed(feed models.Feed) error {
	err := c.repository.Create(feed)
	return err
}
