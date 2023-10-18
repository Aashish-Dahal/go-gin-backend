package models

import "time"

type Feed struct {
	UserID    int       `json:"user_id" validate:"required"`
	User      *User     `json:"user,omitempty"`
	Status    string    `json:"status"`
	Tags      []string  `json:"tags"`
	Images    []string  `json:"images"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// TableName gives table name of model
func (f *Feed) TableName() string {
	return "feeds"
}
