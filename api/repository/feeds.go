package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type FeedRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewFeedRepository(db infrastructure.Database, logger infrastructure.Logger) FeedRepository {
	return FeedRepository{
		db:     db,
		logger: logger,
	}

}

// WithTrx enables repository with transaction
func (c FeedRepository) WithTrx(trxHandle *gorm.DB) FeedRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}
func (c FeedRepository) Create(Feed models.Feed) error {
	return c.db.DB.Create(&Feed).Error
}
