package repository

import (
	"boilerplate-api/dtos"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/paginations"

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

// GetAllFeeds Get All feeds
func (c FeedRepository) GetAllFeeds(pagination paginations.FeedPagination) (feeds []dtos.GetFeedResponse, count int64, err error) {
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Feed{})

	// queryBuilder = queryBuilder.Where("user_id = ?", userId)
	// if pagination.Keyword != "" {
	// 	searchQuery := "%" + pagination.Keyword + "%"
	// 	queryBuilder.Where(c.db.DB.Where("`feeds`.`name` LIKE ?", searchQuery))
	// }

	return feeds, count, queryBuilder.
		Preload("User").
		Find(&feeds).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error
}

// GetAllFeeds Get All feeds
func (c FeedRepository) GetFeedByID(pagination paginations.FeedPagination,userID string) (feeds []dtos.GetFeedResponse, count int64, err error) {
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Feed{})
    queryBuilder.Where("user_id = ?", userID)

	// queryBuilder = queryBuilder.Where("user_id = ?", userId)
	// if pagination.Keyword != "" {
	// 	searchQuery := "%" + pagination.Keyword + "%"
	// 	queryBuilder.Where(c.db.DB.Where("`feeds`.`name` LIKE ?", searchQuery))
	// }

	return feeds, count, queryBuilder.
		Find(&feeds).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error
}
func (c FeedRepository) DeleteFeed(id string) error {
	return c.db.DB.Where("id = ?",id).Delete(&models.Feed{}).Error
}