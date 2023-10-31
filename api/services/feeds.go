package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/dtos"
	"boilerplate-api/models"
	"boilerplate-api/paginations"
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

// GetAllFeeds to get all the Feeds
func (c FeedService) GetAllFeeds(pagination paginations.FeedPagination) ([]dtos.GetFeedResponse, int64, error) {
	return c.repository.GetAllFeeds(pagination)

}
// GetFeedsWithID Get feeds with ID
func (c FeedService) GetFeedByID(pagination paginations.FeedPagination,userID string) ([]dtos.GetFeedResponse, int64, error) {
	return c.repository.GetFeedByID(pagination,userID)

}
// DeleteFeed to delete the feed
func (c FeedService) DeleteFeed(id string) (error) {
	return c.repository.DeleteFeed(id)

}