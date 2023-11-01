package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type LikeRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewLikeRepository(db infrastructure.Database, logger infrastructure.Logger) LikeRepository {
	return LikeRepository{
		db:     db,
		logger: logger,
	}

}

// WithTrx enables repository with transaction
func (c LikeRepository) WithTrx(trxHandle *gorm.DB) LikeRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}
func (c LikeRepository) Create(Like models.Like) error {
	return c.db.DB.Create(&Like).Error
}
