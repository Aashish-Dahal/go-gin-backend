package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewCommentRepository(db infrastructure.Database, logger infrastructure.Logger) CommentRepository {
	return CommentRepository{
		db:     db,
		logger: logger,
	}

}

// WithTrx enables repository with transaction
func (c CommentRepository) WithTrx(trxHandle *gorm.DB) CommentRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

func (c CommentRepository) Create(Comment models.Comment) error {
	return c.db.DB.Create(&Comment).Error
}
