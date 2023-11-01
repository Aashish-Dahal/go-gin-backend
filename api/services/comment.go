package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type CommentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) CommentService {
	return CommentService{
		repository: repository,
	}
}

// CreateFeed to create the Feed
func (c CommentService) Create(Comment models.Comment) error {
	err := c.repository.Create(Comment)
	return err
}
