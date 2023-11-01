package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	UserID    int            `json:"user_id" validate:"required"`
	User      *User          `json:"user,omitempty"`
	FeedID    int            `json:"feed_id" validate:"required"`
	Comment   string         `json:"comment" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName gives table name of model
func (f *Comment) TableName() string {
	return "comments"
}
